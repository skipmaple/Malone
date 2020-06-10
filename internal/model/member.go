// Copyright © 2020. Drew Lee. All rights reserved.

// Package model provides model application from mvc.
package model

import "github.com/jinzhu/gorm"

const (
	Male          = "M"
	Female        = "F"
	UnknownGender = "U"
)

type Member struct {
	gorm.Model
	PhoneNum string `gorm:"varchar(20)" json:"phone_num" form:"phone_num"`
	Password string `gorm:"varchar(40)" json:"-" form:"password"` // 用户密码 md5(password + salt)
	Avatar   string `gorm:"varchar(150)" json:"avatar" form:"avatar"`
	Gender   string `gorm:"varchar(2)" json:"gender" form:"gender"`
	Nickname string `gorm:"varchar(20)" json:"nickname" form:"nickname"`
	Salt     string `gorm:"varchar(10)" json:"-"`
	Online   int    `gorm:"int(10)" json:"online"`
	Token    string `gorm:"varchar(40)" json:"-"`
	Memo     string `gorm:"varchar(150)" json:"memo" form:"memo"`
}
