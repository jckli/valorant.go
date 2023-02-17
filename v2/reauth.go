package val

import (
	"net/http"
	"net"
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

func Reauthenticate(auth *AuthBody) (*http.Response, error) {
	client := &http.Client{Transport: &http.Transport{DialTLS: reauthDialTLS}}
	req, err := http.NewRequest(http.MethodGet, "https://auth.riotgames.com/authorize?redirect_uri=https%3A%2F%2Fplayvalorant.com%2Fopt_in&client_id=play-valorant-web-prod&response_type=token%20id_token&nonce=1", nil)
	if err != nil {
		return nil, err
	}
	reauthHeaders = reauthDefaultHeaders.Clone()
	reauthHeaders.Set("Cookie", auth.Cookies)
	reauthHeaders.Set("Referer", req.URL.Host)
	req.Header = reauthHeaders.Clone()
	fmt.Println(req.Header)
	return client.Do(req)
}
