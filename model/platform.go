package model

type AccountPlatform struct {
	Id            int    `json:"id"`
	Uid           int    `json:"uid"`
	PlatformId    string `json:"platform_id"`
	PlatformToken string `json:"platform_token"`
	Type          int    `json:"type"`
	Nickname      string `json:"nickname"`
	Avatar        string `json:"avatar"`
	CreateAt      int    `json:"create_at"`
	UpdateAt      int    `json:"update_at"`
}
