package valorant

type HandshakeReqBody struct {
	ClientID string `json:"client_id"`
	Nonce int `json:"nonce"`
	RedirectURI string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	Scope string `json:"scope"`
}

type HandshakeRespBody struct {
	Type string `json:"type"`
	Country string `json:"country"`
}

type LoginBody struct {
	Type string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRespBody struct {
	Type string `json:"type"`
	Response struct {
		Mode string `json:"mode"`
		Parameters struct {
			Uri string `json:"uri"`
		} `json:"parameters"`
	} `json:"response"`
	Country string `json:"country"`
}

type ParsedUri struct {
	AccessToken string `json:"access_token"`
	IdToken string `json:"id_token"`
	ExpiresIn string `json:"expires_in"`
}

type EntitlementsRespBody struct {
	Token string `json:"entitlements_token"`
}

type UserInfoRespBody struct {
	Country string `json:"country"`
	UserId string `json:"sub"`
	EmailVerified bool `json:"email_verified"`
	PlayerPlocale string `json:"player_plocale"`
	CountryAt int `json:"country_at"`
	Pw struct {
		CngAt int `json:"cng_at"`
		Reset bool `json:"reset"`
		MustReset bool `json:"must_reset"`
	}
	PhoneNumberVerified bool `json:"phone_number_verified"`
	AccountVerified bool `json:"account_verified"`
	Ppid string `json:"ppid"`
	PlayerLocale string `json:"player_locale"`
	Acct struct {
		Type int `json:"type"`
		State string `json:"state"`
		Adm bool `json:"adm"`
		GameName string `json:"game_name"`
		TagLine string `json:"tag_line"`
		CreatedAt int `json:"created_at"`
	}
	Age int `json:"age"`
	Jti string `json:"jti"`
	Affinity struct {
		Pp string `json:"pp"`
	}
}