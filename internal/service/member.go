// Copyright © 2020. Drew Lee. All rights reserved.

package service

import (
	"KarlMalone/internal/models"
)

type Member struct {
	ID       int64  `form:"id" json:"id" xml:"id"`
	PhoneNum string `form:"phone_num" json:"phone_num" xml:"phone_num"`
	Password string `form:"password" json:"password" xml:"password"`
	Avatar   string `form:"avatar" json:"avatar" xml:"avatar"`
	Gender   string `form:"gender" json:"gender" xml:"gender"`
	Nickname string `form:"nickname" json:"nickname" xml:"nickname"`
	Email    string `form:"email" json:"email" xml:"email"`
	Account  string `form:"account" json:"account" xml:"account"`
	Salt     string
	Online   int
	Token    string
	Memo     string `form:"memo" json:"memo" xml:"memo"`

	CreatedAt int
	UpdatedAt int
	DeletedAt int
}

// member register
func (m *Member) Register() (models.Member, error) {
	data := map[string]interface{}{
		"phone_num": m.PhoneNum,
		"plain_pwd": m.Password,
		"nickname":  m.Nickname,
		"email":     m.Email,
		"account":   m.Account,
		"avatar":    m.Avatar,
		"gender":    m.Gender,
		"memo":      m.Memo,
	}

	member, err := models.RegisterMember(data)

	return member, err
}

// member login
func (m *Member) Login() (models.Member, error) {
	return models.LoginMember(map[string]interface{}{
		"account":   m.Account,
		"plain_pwd": m.Password,
	})
}

// find member by ID
func (m *Member) Find() models.Member {
	return models.FindMember(m.ID)
}

// search member by phone_num
func (m *Member) FindByPhoneNum() models.Member {
	return models.FindMemberByPhoneNum(m.PhoneNum)
}

// search member by nickname
func (m *Member) FindByNickname() models.Member {
	return models.FindMemberByNickname(m.Nickname)
}

// search member by email
func (m *Member) FindByEmail() models.Member {
	return models.FindMemberByEmail(m.Email)
}
