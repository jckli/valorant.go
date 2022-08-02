package valorant

import (
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"
	"net"
	"strings"
	"net/url"
	tls "github.com/refraction-networking/utls"
)

var (
	defaultHeaders = http.Header{
		"Content-Type": {"application/json"},
		"Cookie":       {""},
		"User-Agent":   {"RiotClient/43.0.1.4195386.4190634 rso-auth (Windows; 10;;Professional, x64)"},
	}
	authHeaders = http.Header{}
	tlsConfig   = tls.Config{
		CipherSuites: []uint16{tls.TLS_AES_128_GCM_SHA256},
		MinVersion:   tls.VersionTLS13,
	}
)

// Custom TLS dialer - Credit to fyraux (https://github.com/fyraux/go-rso)
func dialTLS(network, addr string) (net.Conn, error) {
	netConn, err := net.Dial(network, addr)
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
	if err = tlsConn.Handshake(); err != nil {
		return nil, err
	}

	return tlsConn, nil
}

func httpRequest (method, url string, body interface{}) (*http.Response, error) {
	client := &http.Client{Transport: &http.Transport{DialTLS: dialTLS}}
	json_data, _ := json.Marshal(body)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	req.Header = authHeaders.Clone()
	req.Header.Set("Referer", req.URL.Host)
	return client.Do(req)
}

func parseCookies(cookies []string, subs string) (string, error) {
	for _, cookie := range cookies {
		if strings.Contains(cookie, subs) {
			return cookie, nil
		}
	}
	return "", fmt.Errorf("could not find %s", subs)
}

func getTokens(uri string) (*ParsedUri, error) {
	parsedUrl, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	query, err := url.ParseQuery(parsedUrl.Fragment)
	if err != nil {
		return nil, err
	}
	access_token := query.Get("access_token")
	id_token := query.Get("id_token")
	expires_in := query.Get("expires_in")

	return &ParsedUri{
		AccessToken: access_token,
		IdToken:     id_token,
		ExpiresIn:   expires_in,
	}, nil
}

func Handshake() (string, error) {
	resp, err := httpRequest(http.MethodPost, 
		"https://auth.riotgames.com/api/v1/authorization", 
		HandshakeReqBody{
			ClientID:     "play-valorant-web-prod",
			Nonce:        1,
			RedirectURI:  "https://playvalorant.com/opt_in/",
			ResponseType: "token id_token",
			Scope:        "account openid",
		},
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body := new(HandshakeRespBody)
	json.NewDecoder(resp.Body).Decode(body)
	cookie, err := parseCookies(resp.Header["Set-Cookie"], "asid")
	if err != nil {
		return "", err
	}

	return cookie, nil
}

func Login(username, password string) (*ParsedUri, string, error) {
	resp, err := httpRequest(http.MethodPut,
		"https://auth.riotgames.com/api/v1/authorization",
		LoginBody{
			Type:        "auth",
			Username:    username,
			Password:    password,
		},
	)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	body := new(LoginRespBody)
	json.NewDecoder(resp.Body).Decode(body)

	parsedUri, err := getTokens(body.Response.Parameters.Uri)
	if err != nil {
		return nil, "", err
	}

	cookie, err := parseCookies(resp.Header["Set-Cookie"], "ssid")
	if err != nil {
		return nil, "", err
	}

	return parsedUri, cookie, nil
}

func getEntitlements() (string, error) {
	resp, err := httpRequest(http.MethodPost,
		"https://entitlements.auth.riotgames.com/api/token/v1",
		nil,
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body := new(EntitlementsRespBody)
	json.NewDecoder(resp.Body).Decode(body)
	token := body.Token

	return token, nil
}

func getUserInfo() (string, error) {
	resp, err := httpRequest(http.MethodPost,
		"https://auth.riotgames.com/userinfo",
		nil,
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body := new(UserInfoRespBody)
	json.NewDecoder(resp.Body).Decode(body)

	return body.UserId, nil
}

func Authentication(username, password string) (string, http.Header) {
	authHeaders = defaultHeaders.Clone()
	cookie, err := Handshake()
	if err != nil {
		return "", nil
	}
	authHeaders.Set("Cookie", cookie)
	parsedUri, cookie, err := Login(username, password)
	if err != nil {
		return "", nil
	}
	authHeaders.Set("Cookie", cookie)
	authHeaders.Set("Authorization", "Bearer " + parsedUri.AccessToken)
	token, err := getEntitlements()
	if err != nil {
		return "", nil
	}
	authHeaders.Set("X-Riot-Entitlements-JWT", token)
	userId, err := getUserInfo()
	if err != nil {
		return "", nil
	}

	return userId, authHeaders
}