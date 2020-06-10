// Copyright © 2020. Drew Lee. All rights reserved.

// Package args provides basic args.
package args

import (
	"fmt"
	"time"
)

type PageArg struct {
	PageFrom int       `json:"page_from" form:"page_from"` // 从哪页开始
	PageSize int       `json:"page_size" form:"page_size"` // 每页大小
	KeyWord  string    `json:"key_word" form:"key_word"`   // 关键词
	Asc      string    `json:"asc" form:"asc"`
	Desc     string    `json:"desc" form:"desc"`
	Name     string    `json:"name" form:"name"`
	MemberId uint      `json:"member_id" form:"member_id"`
	DstId    uint      `json:"dst_id" form:"dst_id"`
	DateFrom time.Time `json:"date_from" form:"date_from"` // 时间点1
	DateTo   time.Time `json:"date_to" form:"date_to"`     // 时间点2
	Total    int64     `json:"total" form:"total"`
}

func (p *PageArg) GetPageSize() int {
	if p.PageSize == 0 {
		return 100
	} else {
		return p.PageSize
	}
}

func (p *PageArg) GetPageFrom() int {
	if p.PageFrom < 0 {
		return 0
	} else {
		return p.PageFrom
	}
}

func (p *PageArg) GetOrderBy() string {
	if len(p.Asc) > 0 {
		return fmt.Sprintf("%s asc", p.Asc)
	} else if len(p.Desc) > 0 {
		return fmt.Sprintf("%s desc", p.Desc)
	} else {
		return ""
	}
}
