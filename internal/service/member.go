// Copyright © 2020. Drew Lee. All rights reserved.

package service

import (
	"KarlMalone/internal/db"
	"KarlMalone/internal/model"
	"KarlMalone/pkg/logger"
	"KarlMalone/pkg/util"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/jinzhu/gorm"

	"go.uber.org/zap"
)

type MemberService struct{}

// member register
func (s *MemberService) Register(phoneNum, plainPwd, nickname, avatar, gender string) (model.Member, error) {
	member := model.Member{}
	if err := db.Orm.Where("phone_num = ?", phoneNum).Take(&member).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		recordErrLog("member register service find member error", member, err)
		return member, err
	}

	if member.ID > 0 {
		err := errors.New("phone number has been registered")
		recordErrLog("member register service member phone_num error", member, err)
		return member, err
	}

	member.PhoneNum = phoneNum
	member.Avatar = avatar
	member.Gender = gender
	member.Nickname = nickname
	member.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	member.Password = util.MakePwd(plainPwd, member.Salt)

	if err := db.Orm.Create(&member).Error; err != nil {
		recordErrLog("member register service create member error", member, err)
		return member, err
	}

	return member, nil
}

// member login
func (s *MemberService) Login(phoneNum, plainPwd string) (model.Member, error) {
	member := model.Member{}
	if err := db.Orm.Where("phone_num = ?", phoneNum).First(&member).Error; err != nil {
		recordErrLog("member login service find member error", member, err)
		return member, err
	}
	if member.ID == 0 {
		err := errors.New("member not exist")
		return member, err
	}

	if !util.ValidatePwd(plainPwd, member.Salt, member.Password) {
		//logger.Debug("member login service validate password not equal", zap.Any("plainPwd", plainPwd), zap.String("salt", member.Salt), zap.String("password", member.Password))
		return member, errors.New("password wrong")
	}

	// refresh member's login token
	member.Token = util.GenRandomStr(32)
	if err := db.Orm.Save(&member).Error; err != nil {
		recordErrLog("member login service update member error", member, err)
		return member, err
	}

	return member, nil
}

// find member by ID
func (s *MemberService) Find(memberId uint) (model.Member, error) {
	member := model.Member{}
	if err := db.Orm.First(&member, memberId).Error; err != nil {
		logger.Error("member find service find member error", zap.Uint("memberId", memberId), zap.String("reason", err.Error()))
		return model.Member{}, err
	}

	return member, nil
}

func recordErrLog(msg string, member model.Member, err error) {
	jMember, _ := json.Marshal(member)
	logger.Error(msg, zap.String("member", string(jMember)), zap.String("reason", err.Error()))
}
