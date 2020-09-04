// Copyright © 2020. Drew Lee. All rights reserved.

// Package models provides models application from mvc.
package models

import (
	"KarlMalone/pkg/logger"
	"KarlMalone/pkg/util"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"go.uber.org/zap"

	"github.com/jinzhu/gorm"
)

const (
	Male          = "M"
	Female        = "F"
	UnknownGender = "U"
)

type Member struct {
	Model
	PhoneNum string `gorm:"varchar(20)" json:"phone_num" form:"phone_num"`
	Password string `gorm:"varchar(40)" json:"-" form:"password"` // validate md5(password + salt)
	Avatar   string `gorm:"varchar(150)" json:"avatar" form:"avatar"`
	Gender   string `gorm:"varchar(2)" json:"gender" form:"gender"`
	Nickname string `gorm:"varchar(20)" json:"nickname" form:"nickname"`
	Email    string `gorm:"varchar(20)" json:"email" form:"email"`
	Salt     string `gorm:"varchar(10)" json:"-"`
	Online   int    `gorm:"int(10)" json:"online"` // 0-offline 1-online
	Token    string `gorm:"varchar(150)" json:"token"`
	Memo     string `gorm:"varchar(150)" json:"memo" form:"memo"`
}

// RegisterMember register member with params data
func RegisterMember(data map[string]interface{}) (Member, error) {
	member := Member{}
	if err := db.Where("phone_num = ?", data["phone_num"].(string)).Take(&member).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		logger.Error("model member register find error", zap.String("phone_num", data["phone_num"].(string)), zap.String("error", err.Error()))
		return Member{}, err
	}

	if member.ID > 0 {
		err := errors.New("phone number has been registered")
		logger.Error("model member register phone_num has been registered error", zap.String("phone_num", data["phone_num"].(string)), zap.String("error", err.Error()))
		return Member{}, err
	}

	member.PhoneNum = data["phone_num"].(string)
	member.Avatar = data["avatar"].(string)
	member.Gender = data["gender"].(string)
	member.Nickname = data["nickname"].(string)
	member.Email = data["email"].(string)
	member.Memo = data["memo"].(string)
	member.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	member.Password = util.MakePwd(data["plain_pwd"].(string), member.Salt)

	if err := db.Create(&member).Error; err != nil {
		logger.Error("model member register create error", zap.Any("member", member), zap.String("error", err.Error()))
		return Member{}, err
	}

	return member, nil
}

func LoginMember(data map[string]interface{}) (Member, error) {
	member := Member{}

	accountName := ""
	if strings.HasSuffix(data["account"].(string), ".com") { // !!! can strengthen
		accountName = "email"
	} else {
		accountName = "phone_num"
	}
	if err := db.Where(fmt.Sprintf("%s = ?", accountName), data["account"].(string)).First(&member).Error; err != nil {
		return Member{}, err
	}

	if member.ID == 0 {
		return Member{}, errors.New("member not exist")
	}
	if !util.ValidatePwd(data["plain_pwd"].(string), member.Salt, member.Password) {
		return Member{}, errors.New("password wrong")
	}

	// 刷新token
	// token过期时间为3小时，客户端需要保存token，token过期之前客户端主动请求刷新
	member.Token, _ = util.GenerateToken(member.ID)
	// 更新在线状态
	member.Online = 1

	if err := db.Save(&member).Error; err != nil {
		return Member{}, err
	}

	return member, nil
}

func LogoutMember(memberId int64) bool {
	member := Member{}
	db.First(&member, memberId)
	// member 不存在 或 处于离线状态
	if member.ID == 0 || member.Online == 0 {
		return false
	}

	member.Online = 0
	member.Token = ""
	if err := db.Save(&member).Error; err != nil {
		return false
	}

	return true
}

// FindMember provides find member by id
func FindMember(memberId int64) (Member, error) {
	member := Member{}
	db.First(&member, memberId)
	logger.Debug("find member", zap.Int64("query member id", memberId), zap.Int64("res member id", member.ID))
	if member.ID == 0 {
		return Member{}, errors.New("member not exist")
	}
	return member, nil
}

// FindMemberByPhoneNum provides member by phone_num
func FindMemberByPhoneNum(phoneNum string) Member {
	member := Member{}
	db.Where("phone_num = ?", phoneNum).Take(&member)
	return member
}

// FindMemberByNickname provides member by nickname
func FindMemberByNickname(nickname string) Member {
	member := Member{}
	db.Where("nickname = ?", nickname).Take(&member)
	return member
}

// FindMemberByNickname provides member by nickname
func FindMemberByEmail(email string) Member {
	member := Member{}
	db.Where("email = ?", email).Take(&member)
	return member
}
