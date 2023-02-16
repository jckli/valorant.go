package val

type HandshakeReqBody struct {
	ClientID string `json:"client_id"`
	Nonce int `json:"nonce"`
	RedirectURI string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	Scope string `json:"scope"`
}

type HandshakeResp struct {
	Type string `json:"type"`
	Country string `json:"country"`
}

type LoginReqBody struct {
	Type string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Type string `json:"type"`
	Response struct {
		Mode string `json:"mode"`
		Parameters struct {
			Uri string `json:"uri"`
		} `json:"parameters"`
	} `json:"response"`
	Country string `json:"country"`
}

type ParsedUriResp struct {
	AccessToken string `json:"access_token"`
	IdToken string `json:"id_token"`
	ExpiresIn string `json:"expires_in"`
}

type EntitlementsResp struct {
	Token string `json:"entitlements_token"`
}


type RegionReqBody struct {
	IdToken string `json:"id_token"`
}

type RegionResp struct {
	PasToken string `json:"token"`
	Affinities struct {
		Pbe string `json:"pbe"`
		Live string `json:"live"`
	} `json:"affinities"`
}

type ClientVersionResp struct {
	Status string `json:"status"`
	Data struct {
		ManifestId string `json:"manifest_id"`
		Branch string `json:"branch"`
		Version string `json:"version"`
		BuildVersion string `json:"buildVersion"`
		EngineVersion string `json:"engineVersion"`
		RiotClientVersion string `json:"riotClientVersion"`
		BuildData string `json:"buildData"`
	} `json:"data"`
}

type UserInfoResp struct {
	Country string `json:"country"`
	UserId string `json:"sub"`
	EmailVerified bool `json:"email_verified"`
	PlayerPlocale string `json:"player_plocale"`
	CountryAt int `json:"country_at"`
	Pw struct {
		CngAt int `json:"cng_at"`
		Reset bool `json:"reset"`
		MustReset bool `json:"must_reset"`
	} `json:"pw"`
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
	} `json:"acct"`
	Age int `json:"age"`
	Jti string `json:"jti"`
	Affinity struct {
		Pp string `json:"pp"`
	} `json:"affinity"`
}

type AuthBody struct {
	Puuid string 
	Cookies string
	Region string
	AccessToken string
	Token string
	Version string
}
