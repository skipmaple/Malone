// Copyright © 2020. Drew Lee. All rights reserved.

package controller

import (
	"KarlMalone/internal/args"
	"KarlMalone/internal/model"
	"KarlMalone/internal/service"
	"KarlMalone/pkg/logger"
	"KarlMalone/pkg/util"
	"net/http"

	"go.uber.org/zap"
)

var contactService service.ContactService

// add friend
func AddFriend(w http.ResponseWriter, r *http.Request) {
	var arg args.AddNewMember
	if err := util.Bind(r, &arg); err != nil {
		logger.Error("contact controller bind error", zap.String("reason", err.Error()))
	}

	friend := contactService.SearchMemberByNickname(arg.DstName)
	if friend.ID == 0 {
		util.RespFail(w, "friend you add not exist")
	} else {
		err := contactService.AddFriend(arg.MemberId, friend.ID)
		if err != nil {
			util.RespFail(w, err.Error())
		} else {
			util.RespOk(w, nil, "add friend success")
		}
	}
}

// load friend list
func LoadFriend(w http.ResponseWriter, r *http.Request) {
	arg := args.ContactArg{}
	if err := util.Bind(r, &arg); err != nil {
		logger.Error("contact controller bind error", zap.String("reason", err.Error()))
	}
	members := contactService.SearchFriend(arg.MemberId)
	util.RespOkList(w, members, len(members))
}

// create group
func CreateGroup(w http.ResponseWriter, r *http.Request) {
	group := model.Group{}
	if err := util.Bind(r, &group); err != nil {
		logger.Error("contact controller bind error", zap.String("reason", err.Error()))
	}
	group, err := contactService.CreateGroup(group)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, group, "")
	}
}

// join group
func JoinGroup(w http.ResponseWriter, r *http.Request) {
	arg := args.AddNewMember{}
	if err := util.Bind(r, &arg); err != nil {
		logger.Error("contact controller bind error", zap.String("reason", err.Error()))
	}
	group := contactService.SearchGroupByName(arg.DstName)
	if group.ID == 0 {
		util.RespFail(w, "group you want add not exist")
	} else {
		err := contactService.JoinGroup(arg.MemberId, group.ID)
		if err != nil {
			util.RespFail(w, err.Error())
		} else {
			// refresh member's group msg
			AddGroupId(arg.MemberId, group.ID)
			util.RespOk(w, nil, "")
		}
	}
}

// load group list
func LoadGroup(w http.ResponseWriter, r *http.Request) {
	arg := args.ContactArg{}
	if err := util.Bind(r, &arg); err != nil {
		logger.Error("contact controller bind error", zap.String("reason", err.Error()))
	}
	groups := contactService.SearchGroup(arg.MemberId)
	util.RespOkList(w, groups, len(groups))
}

//func AddGroupId(memberId uint, groupId uint) {
//
//}
