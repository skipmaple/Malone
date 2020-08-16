// Copyright © 2020. Drew Lee. All rights reserved.

package service

import (
	"KarlMalone/internal/models"
)

type Group struct {
	ID      int64  `form:"id" json:"id"`
	Name    string `form:"name" json:"name"`
	OwnerId int64  `form:"owner_id" json:"owner_id"`
	Icon    string `form:"icon" json:"icon"`
	Cat     int    `form:"cat" json:"cat"`
	Memo    string `form:"memo" json:"memo"`
}

// search group by group_name
func (g *Group) FindGroupByName() models.Group {
	return models.FindGroupByName(g.Name)
}

// create group
func (g *Group) CreateGroup() (resGroup models.Group, err error) {
	data := map[string]interface{}{
		"name":     g.Name,
		"owner_id": g.OwnerId,
		"icon":     g.Icon,
		"memo":     g.Memo,
	}
	return models.CreateGroup(data)
}

// find members by groupId
func (g *Group) FindGroupMembersByGroupId() []models.Member {
	return models.FindGroupMembersByGroupId(g.ID)
}
