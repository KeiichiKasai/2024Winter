package model

import "time"

type UserInfo struct {
	Uid      int       `form:"uid,omitempty"`
	Username string    `form:"username,omitempty"`
	Password string    `form:"password,omitempty"`
	Phone    string    `form:"phone,omitempty"`
	Address  string    `form:"address,omitempty"`
	Gender   int       `form:"gender,omitempty"`
	Birthday time.Time `form:"birthday,omitempty"`
}
