// Copyright © 2020. Drew Lee. All rights reserved.

package controller

import (
	"KarlMalone/internal/model"
	"KarlMalone/internal/service"
	"KarlMalone/pkg/logger"
	"KarlMalone/pkg/util"
	"net/http"

	"go.uber.org/zap"
)

var MemberService service.MemberService

// member register
func MemberRegister(w http.ResponseWriter, r *http.Request) {
	//_ = r.ParseForm()
	member := model.Member{}
	if err := util.Bind(r, &member); err != nil {
		logger.Error("member register controller bind error", zap.String("reason", err.Error()))
	}

	member, err := MemberService.Register(member.PhoneNum, member.Password, member.Nickname, member.Avatar, member.Gender)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, member, "")
	}
}

// member login
func MemberLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logger.Error("controller member login parseForm error", zap.String("reason", err.Error()))
	}

	phoneNum := r.PostForm.Get("phone_num")
	plainPwd := r.PostForm.Get("password")

	// validate parameters
	if len(phoneNum) == 0 || len(plainPwd) == 0 {
		util.RespFail(w, "username or password error")
	}

	member, err := MemberService.Login(phoneNum, plainPwd)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, member, "")
	}
}
