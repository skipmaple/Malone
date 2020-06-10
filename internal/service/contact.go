// Copyright © 2020. Drew Lee. All rights reserved.

package service

import (
	"KarlMalone/internal/db"
	"KarlMalone/internal/model"
	"errors"
)

type ContactService struct{}

// add friend
func (c *ContactService) AddFriend(memberId uint, dstId uint) error {
	if dstId == memberId {
		return errors.New("can't add yourself as friend")
	}

	friend := model.Contact{}
	db.Orm.Where("owner_id = ? AND dst_obj = ? And cat = ?", memberId, dstId, model.CatIndividual).Take(&friend)
	if friend.ID > 0 {
		return errors.New("friends have been added before")
	}

	// start transaction
	tx := db.Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	contact1 := model.Contact{
		OwnerId: memberId,
		DstId:   dstId,
		Cat:     model.CatIndividual,
	}
	contact2 := model.Contact{
		OwnerId: dstId,
		DstId:   memberId,
		Cat:     model.CatIndividual,
	}
	if err := tx.Create(contact1).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(contact2).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// search group by owner_id
func (c *ContactService) SearchGroup(memberId uint) []model.Group {
	contacts := make([]model.Contact, 0)
	groupIds := make([]uint, 0)

	db.Orm.Where("owner_id = ? AND cat = ?", memberId, model.CatGroup).Find(&contacts)
	for _, contact := range contacts {
		groupIds = append(groupIds, contact.DstId)
	}
	groups := make([]model.Group, 0)
	if len(groupIds) == 0 {
		return groups
	}
	db.Orm.Where("id IN (?)", groupIds).Find(&groups)
	return groups
}

// search group by group_name
func (c *ContactService) SearchGroupByName(name string) model.Group {
	group := model.Group{}
	db.Orm.Where("name = ?", name).Take(&group)
	return group
}

// search member's friends by member_id
func (c *ContactService) SearchFriend(memberId uint) []model.Member {
	contacts := make([]model.Contact, 0)
	friendMemberIds := make([]uint, 0)
	db.Orm.Where("owner_id = ? AND cat = ?", memberId, model.CatIndividual).Find(&contacts)
	for _, friend := range contacts {
		friendMemberIds = append(friendMemberIds, friend.DstId)
	}
	members := make([]model.Member, 0)
	if len(friendMemberIds) == 0 {
		return members
	}
	db.Orm.Where("id IN (?)", friendMemberIds).Find(&members)
	return members
}

// search member by phone_num
func (c *ContactService) SearchMemberByPhoneNum(phoneNum string) model.Member {
	member := model.Member{}
	db.Orm.Where("phone_num = ?", phoneNum).Take(&member)
	return member
}

// search member by nickname
func (c *ContactService) SearchMemberByNickname(nickname string) model.Member {
	member := model.Member{}
	db.Orm.Where("nickname = ?", nickname).Take(&member)
	return member
}

// search group_ids by owner_id
func (c *ContactService) SearchGroupIds(memberId uint) []uint {
	contacts := make([]model.Contact, 0)
	groupIds := make([]uint, 0)

	db.Orm.Where("owner_id = ? AND cat = ?", memberId, model.CatGroup).Find(&contacts)
	for _, contact := range contacts {
		groupIds = append(groupIds, contact.DstId)
	}
	return groupIds
}

// create group
func (c *ContactService) CreateGroup(group model.Group) (resGroup model.Group, err error) {
	if len(group.Name) == 0 {
		err = errors.New("no group name error")
		return resGroup, err
	}
	if group.OwnerId == 0 {
		err = errors.New("please sign in first")
		return resGroup, err
	}

	var groupCount uint
	err = db.Orm.Where(&model.Group{OwnerId: group.OwnerId}).Count(&groupCount).Error
	if err != nil {
		return model.Group{}, err
	}
	if groupCount >= 5 {
		err = errors.New("member can create group max count is 5")
		return model.Group{}, err
	}

	tx := db.Orm.Begin()
	if err = tx.Create(&group).Error; err != nil {
		tx.Rollback()
		return model.Group{}, err
	}
	err = db.Orm.Create(&model.Contact{
		OwnerId: group.OwnerId,
		DstId:   group.ID,
		Cat:     model.CatGroup,
	}).Error
	if err != nil {
		tx.Rollback()
		return model.Group{}, err
	}
	tx.Commit()

	return group, nil
}

// member add group
func (c *ContactService) JoinGroup(memberId, groupId uint) error {
	contact := model.Contact{
		OwnerId: memberId,
		DstId:   groupId,
		Cat:     model.CatGroup,
	}
	err := db.Orm.Where(contact).FirstOrCreate(&contact).Error
	return err
}
