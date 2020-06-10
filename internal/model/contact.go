// Copyright © 2020. Drew Lee. All rights reserved.

package model

import "github.com/jinzhu/gorm"

const (
	CatIndividual = 0x01 // 个人
	CatGroup      = 0x02 // 群组
)

// 好友和群存在这个表里边
// 可根据业务做拆分
type Contact struct {
	gorm.Model
	OwnerId uint   `gorm:"bigint(20)" json:"owner_id"`
	DstId   uint   `gorm:"bigint(20)" json:"dst_id"` // 对端信息是谁的
	Cat     int    `gorm:"int(11)" json:"cat"`
	Memo    string `gorm:"varchar(150)" json:"memo"`
}
