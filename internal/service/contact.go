// Copyright © 2020. Drew Lee. All rights reserved.

package service

import (
	"KarlMalone/internal/models"
)

type Contact struct {
	ID      int64  `form:"id" json:"id" xml:"id"`
	OwnerId int64  `gorm:"bigint(20)" json:"owner_id"`
	DstId   int64  `gorm:"bigint(20)" json:"dst_id"`
	Cat     int    `gorm:"int(11)" json:"cat"`
	Memo    string `gorm:"varchar(150)" json:"memo"`

	CreatedAt int
	UpdatedAt int
	DeletedAt int
}

// add friend
func (c *Contact) AddFriend() error {
	data := map[string]interface{}{
		"owner_id": c.OwnerId,
		"dst_id":   c.DstId,
	}

	if err := models.AddFriend(data); err != nil {
		return err
	}

	return nil
}

// member join group
func (c *Contact) JoinGroup() error {
	data := map[string]interface{}{
		"owner_id": c.OwnerId,
		"dst_id":   c.DstId,
	}
	return models.JoinGroup(data)
}

// find all friends by owner_id
func (c *Contact) FindAllFriend() []models.Member {
	return models.FindAllFriend(c.OwnerId)
}

// search groups by owner_id which member has joined
func (c *Contact) FindGroup() []models.Group {
	return models.FindGroup(c.OwnerId)
}

// search group_ids by owner_id
func (c *Contact) FindGroupIds() []int64 {
	return models.FindGroupIds(c.OwnerId)
}
