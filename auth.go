package valorant

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	tls "github.com/refraction-networking/utls"
	"github.com/valyala/fasthttp"
	"net"
	"net/url"
	"strings"
)

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomStringURLSafe(n int) string {
	b, _ := generateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b)
}

var (
	defaultHeaders = map[string]string{
		"Content-Type":  "application/json",
		"Cache-Control": "no-cache",
		"User-Agent":    GenerateRandomStringURLSafe(111),
	}
	authHeaders = map[string]string{}
	tlsConfig   = &tls.Config{
		CipherSuites: []uint16{tls.TLS_CHACHA20_POLY1305_SHA256},
		MinVersion:   tls.VersionTLS13,
	}
)

type Auth struct {
	Client      *fasthttp.Client
	CookieJar   string
	Region      string
	AccessToken string
	IdToken     string
	ExpiresIn   string
	Token       string
	Version     string
	UserInfo    userInfoResp
}

type handshakeBody struct {
	ClientID     string `json:"client_id"`
	Nonce        string `json:"nonce"`
	RedirectURI  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	Scope        string `json:"scope"`
}

type handshakeRespBody struct {
	Type    string `json:"type"`
	Country string `json:"country"`
}

type loginBody struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginRespBody struct {
	Type     string `json:"type"`
	Response struct {
		Mode       string `json:"mode"`
		Parameters struct {
			Uri string `json:"uri"`
		} `json:"parameters"`
	} `json:"response"`
	Country string `json:"country"`
}

type userInfoResp struct {
	Country       string `json:"country"`
	UserId        string `json:"sub"`
	EmailVerified bool   `json:"email_verified"`
	PlayerPlocale string `json:"player_plocale"`
	CountryAt     int    `json:"country_at"`
	Pw            struct {
		CngAt     int  `json:"cng_at"`
		Reset     bool `json:"reset"`
		MustReset bool `json:"must_reset"`
	} `json:"pw"`
	PhoneNumberVerified bool   `json:"phone_number_verified"`
	AccountVerified     bool   `json:"account_verified"`
	Ppid                string `json:"ppid"`
	PlayerLocale        string `json:"player_locale"`
	Acct                struct {
		Type      int    `json:"type"`
		State     string `json:"state"`
		Adm       bool   `json:"adm"`
		GameName  string `json:"game_name"`
		TagLine   string `json:"tag_line"`
		CreatedAt int    `json:"created_at"`
	} `json:"acct"`
	Age      int    `json:"age"`
	Jti      string `json:"jti"`
	Affinity struct {
		Pp string `json:"pp"`
	} `json:"affinity"`
}

type regionReq struct {
	IdToken string `json:"id_token"`
}

type regionResp struct {
	PasToken   string `json:"token"`
	Affinities struct {
		Pbe  string `json:"pbe"`
		Live string `json:"live"`
	} `json:"affinities"`
}

type clientVersionResp struct {
	Status int `json:"status"`
	Data   struct {
		ManifestId        string `json:"manifest_id"`
		Branch            string `json:"branch"`
		Version           string `json:"version"`
		BuildVersion      string `json:"buildVersion"`
		EngineVersion     string `json:"engineVersion"`
		RiotClientVersion string `json:"riotClientVersion"`
		BuildData         string `json:"buildData"`
	} `json:"data"`
}

func dialTLS(addr string) (net.Conn, error) {
	netConn, err := fasthttp.Dial(addr)
	if err != nil {
		return nil, err
	}
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	config := tlsConfig.Clone()
	config.ServerName = host

	tlsConn := tls.UClient(netConn, config, tls.HelloGolang)
	if err := tlsConn.Handshake(); err != nil {
		return nil, err
	}

	return tlsConn, nil
}

func createClient() *fasthttp.Client {
	client := &fasthttp.Client{
		Dial: dialTLS,
	}

	return client
}

func (a *Auth) httpRequest(method, url string, body interface{}) (r []byte, ok bool) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod(method)
	req.SetRequestURI(url)
	for k, v := range authHeaders {
		req.Header.Add(k, v)
	}
	req.Header.AddBytesV("Referer", req.URI().Host())

	if body != nil {
		bodyBytes, _ := json.Marshal(body)
		req.SetBody(bodyBytes)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := a.Client.Do(req, resp)
	if err != nil {
		return nil, false
	}
	r = resp.Body()
	return r, true
}

func parseCookies(cookies []string, subs string) (string, error) {
	for _, cookie := range cookies {
		if strings.Contains(cookie, subs) {
			return cookie, nil
		}
	}
	return "", fmt.Errorf("could not find %s", subs)
}

func (a *Auth) getTokens(uri string) (ok bool) {
	parsedUrl, err := url.Parse(uri)
	if err != nil {
		return false
	}
	query, err := url.ParseQuery(parsedUrl.Fragment)
	if err != nil {
		return false
	}
	access_token := query.Get("access_token")
	id_token := query.Get("id_token")
	expires_in := query.Get("expires_in")

	a.AccessToken = access_token
	a.IdToken = id_token
	a.ExpiresIn = expires_in

	return true
}

func (a *Auth) handshake() (cookie string, ok bool) {
	body := handshakeBody{
		ClientID:     "play-valorant-web-prod",
		Nonce:        GenerateRandomStringURLSafe(16),
		RedirectURI:  "https://playvalorant.com/opt_in",
		ResponseType: "token id_token",
		Scope:        "account openid",
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("POST")
	req.SetRequestURI("https://auth.riotgames.com/api/v1/authorization")
	for k, v := range authHeaders {
		req.Header.Add(k, v)
	}
	req.Header.AddBytesV("Referer", req.URI().Host())

	bodyBytes, _ := json.Marshal(body)
	req.SetBody(bodyBytes)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := a.Client.Do(req, resp)
	if err != nil {
		return "", false
	}

	respBody := handshakeRespBody{}
	if err := json.Unmarshal(resp.Body(), &respBody); err != nil {
		return "", false
	}
	var cookies []string
	resp.Header.VisitAll(func(key, value []byte) {
		if string(key) == "Set-Cookie" {
			cookies = append(cookies, string(value))
		}
	})

	cookie, err = parseCookies(cookies, "asid")
	if err != nil {
		return "", false
	}
	return cookie, true
}

func (a *Auth) login(username, password string) (cookie string, ok bool) {
	body := loginBody{
		Type:     "auth",
		Username: username,
		Password: password,
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("PUT")
	req.SetRequestURI("https://auth.riotgames.com/api/v1/authorization")
	for k, v := range authHeaders {
		req.Header.Add(k, v)
	}
	req.Header.AddBytesV("Referer", req.URI().Host())

	bodyBytes, _ := json.Marshal(body)
	req.SetBody(bodyBytes)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := a.Client.Do(req, resp)
	if err != nil {
		return "", false
	}

	respBody := loginRespBody{}
	if err := json.Unmarshal(resp.Body(), &respBody); err != nil {
		return "", false
	}

	if respBody.Type == "error" {
		return "", false
	}

	uri := respBody.Response.Parameters.Uri
	if ok := a.getTokens(uri); !ok {
		return "", false
	}

	var cookies []string
	resp.Header.VisitAll(func(key, value []byte) {
		if string(key) == "Set-Cookie" {
			cookies = append(cookies, string(value))
		}
	})

	cookie, err = parseCookies(cookies, "ssid")
	if err != nil {
		return "", false
	}
	return cookie, true
}

func (a *Auth) getEntitlements() (ok bool) {
	body, ok := a.httpRequest("POST", "https://entitlements.auth.riotgames.com/api/token/v1", nil)
	if !ok {
		return false
	}
	bodyStruct := struct {
		Token string `json:"entitlements_token"`
	}{}
	if err := json.Unmarshal(body, &bodyStruct); err != nil {
		return false
	}

	a.Token = bodyStruct.Token
	return true
}

func (a *Auth) getUser() (ok bool) {
	body, ok := a.httpRequest("POST", "https://auth.riotgames.com/userinfo", nil)
	if !ok {
		return false
	}

	bodyStruct := userInfoResp{}
	if err := json.Unmarshal(body, &bodyStruct); err != nil {
		return false
	}

	a.UserInfo = bodyStruct
	return true
}

func (a *Auth) getRegion() (ok bool) {
	body, ok := a.httpRequest("PUT", "https://riot-geo.pas.si.riotgames.com/pas/v1/product/valorant", regionReq{
		IdToken: a.IdToken,
	})
	if !ok {
		return false
	}

	bodyStruct := regionResp{}
	if err := json.Unmarshal(body, &bodyStruct); err != nil {
		return false
	}

	a.Region = bodyStruct.Affinities.Live

	return true
}

func (a *Auth) getClientVersion() (ok bool) {
	body, ok := a.httpRequest("GET", "https://valorant-api.com/v1/version", nil)
	if !ok {
		return false
	}

	bodyStruct := clientVersionResp{}
	if err := json.Unmarshal(body, &bodyStruct); err != nil {
		return false
	}

	a.Version = bodyStruct.Data.RiotClientVersion

	return true
}

// Login to the Riot Games API
func New(username, password string) (*Auth, error) {
	client := createClient()
	auth := &Auth{
		Client: client,
	}

	authHeaders = defaultHeaders

	cookie, ok := auth.handshake()
	if !ok {
		return nil, fmt.Errorf("could not handshake")
	}
	authHeaders["Cookie"] = cookie

	cookie, ok = auth.login(username, password)
	if !ok {
		return nil, fmt.Errorf("could not login")
	}
	delete(authHeaders, "Cookie")
	authHeaders["Authorization"] = fmt.Sprintf("Bearer %s", auth.AccessToken)

	if ok := auth.getRegion(); !ok {
		return nil, fmt.Errorf("could not get region")
	}

	authHeaders["Cookie"] = cookie
	auth.CookieJar = cookie

	if ok := auth.getEntitlements(); !ok {
		return nil, fmt.Errorf("could not get entitlements")
	}
	authHeaders["X-Riot-Entitlements-JWT"] = auth.Token

	if ok := auth.getClientVersion(); !ok {
		return nil, fmt.Errorf("could not get client version")
	}
	authHeaders["X-Riot-ClientVersion"] = auth.Version

	if ok := auth.getUser(); !ok {
		return nil, fmt.Errorf("could not get user")
	}

	return auth, nil
}
