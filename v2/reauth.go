package val

import (
	"net/http"
	"net"
	"net/url"
	"fmt"
	tls "github.com/refraction-networking/utls"
)

var (
	reauthDefaultHeaders = http.Header{
		"Content-Type": {"application/json"},
		"Cookie":       {""},
		"User-Agent":   {""},
	}
	reauthHeaders = http.Header{}
	reauthTlsConfig = tls.Config{
		CipherSuites: []uint16{tls.TLS_AES_128_GCM_SHA256},
		MinVersion:   tls.VersionTLS13,
	}
)

func reauthDialTLS(network, addr string) (net.Conn, error) {
	netConn, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}

	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	config := reauthTlsConfig.Clone()
	config.ServerName = host

	tlsConn := tls.UClient(netConn, config, tls.HelloGolang)
	if err = tlsConn.Handshake(); err != nil {
		return nil, err
	}

	return tlsConn, nil
}

func Reauthenticate(auth *AuthBody) (*AuthBody, error) {
	client := &http.Client{
		Transport: &http.Transport{DialTLS: reauthDialTLS},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
            redirectUrl := req.URL.String()
            u, err := url.Parse(redirectUrl)
            if err != nil {
                return err
            }
			fragment, _ := url.QueryUnescape(u.Fragment)
   			values, _ := url.ParseQuery(fragment)
			fmt.Println("old access_token: " + auth.AccessToken)
			auth.AccessToken = values.Get("access_token")
			fmt.Println("new access_token: " + auth.AccessToken)
			auth.Token = values.Get("id_token")
            return http.ErrUseLastResponse
        },
	}
	req, err := http.NewRequest(http.MethodGet, "https://auth.riotgames.com/authorize?redirect_uri=https%3A%2F%2Fplayvalorant.com%2Fopt_in&client_id=play-valorant-web-prod&response_type=token%20id_token&nonce=1", nil)
	if err != nil {
		return nil, err
	}
	reauthHeaders = reauthDefaultHeaders.Clone()
	reauthHeaders.Set("Cookie", auth.Cookies)
	reauthHeaders.Set("Referer", req.URL.Host)
	req.Header = reauthHeaders.Clone()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &AuthBody{
		AccessToken: auth.AccessToken,
		Token: auth.Token,
		Cookies: auth.Cookies,
		Region: auth.Region,
		Version: auth.Version,
		Puuid: auth.Puuid,
	}, nil
}
