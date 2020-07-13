package model

type Member struct {
	Id       int    `json:"id"`
	Uid      int    `json:"uid"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
	Role     int    `json:"role"`
	CreateAt int    `json:"create_at"`
	UpdateAt int    `json:"update_at"`
}
