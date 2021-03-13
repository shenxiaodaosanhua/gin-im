package model

import (
	"time"
)

type UserInfo struct {
	Id uint `json:"-"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Desc string `json:"desc"`
	UserId uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (u *UserInfo) TableName() string {
	return "user_info"
}