// Copyright © 2020. Drew Lee. All rights reserved.

package models

import (
	"errors"
)

const (
	CatIndividual = 0x01
	CatGroup      = 0x02
)

type Contact struct {
	Model
	OwnerId int64  `gorm:"bigint(20)" json:"owner_id"`
	DstId   int64  `gorm:"bigint(20)" json:"dst_id"` // group_id or member_id
	Cat     int    `gorm:"int(11)" json:"cat"`
	Memo    string `gorm:"varchar(150)" json:"memo"`
}

// AddFriend provides member add new friend
func AddFriend(data map[string]interface{}) error {
	dstId := data["dst_id"].(int64)
	ownerId := data["owner_id"].(int64)
	if dstId == ownerId {
		return errors.New("can't add yourself as friend")
	}

	friend := Contact{}
	db.Where("owner_id = ? AND dst_id = ? And cat = ?", ownerId, dstId, CatIndividual).Take(&friend)
	if friend.ID > 0 {
		return errors.New("friends have been added before")
	}

	// start transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	contact1 := Contact{
		OwnerId: ownerId,
		DstId:   dstId,
		Cat:     CatIndividual,
	}
	contact2 := Contact{
		OwnerId: dstId,
		DstId:   ownerId,
		Cat:     CatIndividual,
	}
	if err := tx.Create(&contact1).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&contact2).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// FindAllFriend provides all friends by member_id
func FindAllFriend(ownerId int64) []Member {
	contacts := make([]Contact, 0)
	friendMemberIds := make([]int64, 0)
	db.Where("owner_id = ? AND cat = ?", ownerId, CatIndividual).Find(&contacts)
	for _, friend := range contacts {
		friendMemberIds = append(friendMemberIds, friend.DstId)
	}
	members := make([]Member, 0)
	if len(friendMemberIds) == 0 {
		return members
	}
	db.Where("id IN (?)", friendMemberIds).Find(&members)

	return members
}
