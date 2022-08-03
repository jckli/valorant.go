package valorant

// auth structs
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

// contracts structs
type ContractDefinitionsResp struct {
	Definitions []struct {
		Id string `json:"ID"`
		Item struct {
			ItemTypeId string `json:"ItemTypeID"`
			ItemId string `json:"ItemID"`
		} `json:"Item"`
		RequiredEntitlement struct {
			ItemTypeId string `json:"ItemTypeID"`
			ItemId string `json:"ItemID"`
		} `json:"RequiredEntitlement"`
		ProgressionSchedule struct {
			Name string `json:"Name"`
			ProgressionCurrencyId string `json:"ProgressionCurrencyID"`
			ProgressionDeltaPerLevel []int `json:"ProgressionDeltaPerLevel"`
		} `json:"ProgressionSchedule"`
		RewardSchedule struct {
			Id string `json:"ID"`
			Name string `json:"Name"`
			Prerequisites interface{} `json:"Prerequisites"`
			RewardsPerLevel []struct {
				EntitlementRewards []struct {
					ItemTypeId string `json:"ItemTypeID"`
					ItemId string `json:"ItemID"`
					Amount int `json:"Amount"`
				} `json:"EntitlementRewards"`
				WalletRewards interface{} `json:"WalletRewards"`
				CounterRewards interface{} `json:"CounterRewards"`
			} `json:"RewardsPerLevel"`
		} `json:"RewardSchedule"`
		Sidegrades []struct {
			SidegradeId string `json:"SidegradeID"`
			Options []struct {
				Optionid string `json:"OptionID"`
				Cost struct {
					WalletCosts []struct {
						CurrencyId string `json:"CurrencyID"`
						AmountToDeduct int `json:"AmountToDeduct"`
					} `json:"WalletCosts"`
				} `json:"Cost"`
				Rewards []struct {
					ItemTypeId string `json:"ItemTypeID"`
					ItemId string `json:"ItemID"`
					Amount int `json:"Amount"`
				} `json:"Rewards"`
			} `json:"Options"`
			Prerequisites struct {
				RequiredEntitlements []struct {
					ItemTypeId string `json:"ItemTypeID"`
					ItemId string `json:"ItemID"`
				} `json:"RequiredEntitlements"`
			} `json:"Prerequisites"`
		} `json:"Sidegrades"`
	} `json:"Definitions"`
}

type ContractFetchResp struct {
	Version int `json:"Version"`
	Subject string `json:"Subject"`
	Contracts []struct {
		ContractDefinitionId string `json:"ContractDefinitionID"`
		ContractProgression struct {
			TotalProgressionEarned int `json:"TotalProgressionEarned"`
			HighestRewardedLevel interface{} `json:"HighestRewardedLevel"`
		} `json:"ContractProgression"`
		ProgressionLevelReached int `json:"ProgressionLevelReached"`
		ProgressionTowardsNextLevel int `json:"ProgressionTowardsNextLevel"`
	} `json:"Contracts"`
	ProcessedMatches []struct {
		Id string `json:"ID"`
		StartTime int `json:"StartTime"`
		XpGrants interface{} `json:"XPGrants"`
		RewardGrants interface{} `json:"RewardGrants"`
		MissionDeltas interface{} `json:"MissionDeltas"`
		ContractDeltas interface{} `json:"ContractDeltas"`
		CouldProgressMissions bool `json:"CouldProgressMissions"`
	} `json:"ProcessedMatches"`
	ActiveSpecialContract string `json:"ActiveSpecialContract"`
	Missions []struct {
		Id string `json:"ID"`
		Objectives interface{} `json:"Objectives"`
		Complete bool `json:"Complete"`
		ExpirationTime string `json:"ExpirationTime"`
	} `json:"Missions"`
	MissionsMetadata struct {
		NpeCompleted bool `json:"NPE_Completed"`
		WeeklyRefillTime string `json:"WeeklyRefillTime"`
	} `json:"MissionsMetadata"`
}

// coregame structs
type CoregameFetchPlayerResp struct {
	Subject string `json:"Subject"`
	MatchId string `json:"MatchID"`
	Versioon int `json:"Version"`
}

type CoregameFetchMatchResp struct {
	MatchId string `json:"MatchID"`
	Version int `json:"Version"`
	State string `json:"State"`
	MapId string `json:"MapID"`
	ModeId string `json:"ModeID"`
	ProvisioningFlow string `json:"ProvisioningFlow"`
	GamePodId string `json:"GamePodID"`
	AllMucName string `json:"AllMUCName"`
	TeamMucTame string `json:"TeamMUCTame"`
	TeamVoiceId string `json:"TeamVoiceID"`
	IsReconnectable bool `json:"IsReconnectable"`
	ConnectionDetails struct {
		GameServerHosts []string
		GameServerHost string `json:"GameServerHost"`
		GameServerPort int `json:"GameServerPort"`
		GameServerObfuscatedIp int `json:"GameServerObfuscatedIP"`
		GameClientHash int `json:"GameClientHash"`
		PlayerKey string `json:"PlayerKey"`
	} `json:"ConnectionDetails"`
	PostGameDetails interface{} `json:"PostGameDetails"`
	Players []struct {
		Subject string `json:"Subject"`
		TeamId string `json:"TeamID"`
		CharacterId string `json:"CharacterID"`
		PlayerIdentity struct {
			Subject string `json:"Subject"`
			PlayerCardId string `json:"PlayerCardID"`
			PlayerTitleId string `json:"PlayerTitleID"`
			AccountLevel int `json:"AccountLevel"`
			PreferredLevelBorderId string `json:"PreferredLevelBorderID"`
			Incognito bool `json:"Incognito"`
			HideAccountLevel bool `json:"HideAccountLevel"`
		} `json:"PlayerIdentity"`
		SeasonalBadgeInfo struct {
			SeasonId string `json:"SeasonID"`
			NumberOfWins int `json:"NumberOfWins"`
			WinsByTier interface{} `json:"WinsByTier"`
			Rank int `json:"Rank"`
			LeaderboardRank int `json:"LeaderboardRank"`
		} `json:"SeasonalBadgeInfo"`
		IsCoach bool `json:"IsCoach"`
		IsAssociated bool `json:"IsAssociated"`
	} `json:"Players"`
	MatchmakingData struct{
		QueueId string `json:"QueueID"`
		IsRanked bool `json:"IsRanked"`
	} `json:"MatchmakingData"`
}

type CoregameFetchMatchLoadoutsResp struct {
	Loadouts []struct {
		CharacterId string `json:"CharacterID"`
		Loadout struct {
			Sprays struct {
				SpraySelections []struct {
					SocketId string `json:"SocketID"`
					SprayId string `json:"SprayID"`
					LevelId string `json:"LevelID"`
				} `json:"SpraySelections"`
			} `json:"Sprays"`
			Items interface{} `json:"Items"`
		} `json:"Loadout"`
	} `json:"Loadouts"`
}