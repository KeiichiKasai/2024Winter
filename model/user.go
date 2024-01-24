package model

import "time"

type UserInfo struct {
	Uid      int       `form:"uid,omitempty" gorm:"primaryKey;autoIncrement"`
	Username string    `form:"username,omitempty" gorm:"not null"`
	Password string    `form:"password,omitempty" gorm:"not null"`
	Phone    string    `form:"phone,omitempty" gorm:"not null"`
	Address  string    `form:"address,omitempty" gorm:"default:null"`
	Gender   int       `form:"gender,omitempty" gorm:"default:0"`
	Birthday time.Time `form:"birthday,omitempty"`
}
