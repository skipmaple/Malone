// Copyright © 2020. Drew Lee. All rights reserved.

package models

import "errors"

const (
	GroupCategoryCom = 0x01
)

type Group struct {
	Model
	Name    string `gorm:"varchar(30)" json:"name"`
	OwnerId int64  `gorm:"bigint(20)" json:"owner_id"` // 群主
	Icon    string `gorm:"varchar(250)" json:"icon"`   // 群logo
	Cat     int    `gorm:"int(11)" json:"cat"`
	Memo    string `gorm:"varchar(150)" json:"memo"`
}

// FindGroup provides groups member has joined
func FindGroup(ownerId int64) []Group {
	contacts := make([]Contact, 0)
	groupIds := make([]int64, 0)
	db.Where("owner_id = ? AND cat = ?", ownerId, CatGroup).Find(&contacts)

	for _, contact := range contacts {
		groupIds = append(groupIds, contact.DstId)
	}
	groups := make([]Group, 0)
	if len(groupIds) == 0 {
		return groups
	}
	db.Where("id IN (?)", groupIds).Find(&groups)

	return groups
}

// FindGroupByName provides group by group_name
func FindGroupByName(name string) Group {
	group := Group{}
	db.Where("name = ?", name).Take(&group)
	return group
}

// FindGroupIds provides group_ids by owner_id
func FindGroupIds(ownerId int64) []int64 {
	contacts := make([]Contact, 0)
	groupIds := make([]int64, 0)

	db.Where("owner_id = ? AND cat = ?", ownerId, CatGroup).Find(&contacts)
	for _, contact := range contacts {
		groupIds = append(groupIds, contact.DstId)
	}
	return groupIds
}

// CreateGroup
func CreateGroup(data map[string]interface{}) (group Group, err error) {
	name := data["name"].(string)
	ownerId := data["owner_id"].(int64)
	icon := data["icon"].(string)
	memo := data["memo"].(string)
	cat := GroupCategoryCom

	if len(name) == 0 {
		err = errors.New("no group name error")
		return Group{}, err
	}
	if ownerId == 0 {
		err = errors.New("please sign in first")
		return Group{}, err
	}

	var groupCount int
	err = db.Model(&Group{}).Where("owner_id = ?", ownerId).Count(&groupCount).Error
	if err != nil {
		return Group{}, err
	}
	if groupCount >= 5 {
		err = errors.New("member can create group max count is 5")
		return Group{}, err
	}

	group = Group{
		Name:    name,
		OwnerId: ownerId,
		Icon:    icon,
		Cat:     cat,
		Memo:    memo,
	}

	tx := db.Begin()
	if err = tx.Create(&group).Error; err != nil {
		tx.Rollback()
		return Group{}, err
	}
	err = db.Create(&Contact{
		OwnerId: group.OwnerId,
		DstId:   group.ID,
		Cat:     CatGroup,
	}).Error
	if err != nil {
		tx.Rollback()
		return Group{}, err
	}
	tx.Commit()

	return group, nil
}

// JoinGroup
func JoinGroup(data map[string]interface{}) error {
	contact := Contact{
		OwnerId: data["owner_id"].(int64),
		DstId:   data["dst_id"].(int64),
		Cat:     CatGroup,
	}
	resContact := Contact{}
	db.Where(contact).Take(&resContact)
	if resContact.ID > 0 {
		err := errors.New("you have in the group")
		return err
	}
	err := db.Create(&contact).Error
	return err
}
