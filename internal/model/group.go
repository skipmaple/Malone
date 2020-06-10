// Copyright © 2020. Drew Lee. All rights reserved.

package model

import "github.com/jinzhu/gorm"

const (
	GroupCategoryCom = 0x01
)

type Group struct {
	gorm.Model
	Name    string `gorm:"varchar(30)" json:"name"`
	OwnerId uint   `gorm:"bigint(20)" json:"owner_id"` // 群主
	Icon    string `gorm:"varchar(250)" json:"icon"`   // 群logo
	Cat     int    `gorm:"int(11)" json:"cat"`
	Memo    string `gorm:"varchar(150)" json:"memo"`
}
