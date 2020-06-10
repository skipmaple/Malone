// Copyright © 2020. Drew Lee. All rights reserved.

package args

type ContactArg struct {
	PageArg
	//MemberId uint // 这个参数是不是重复了？？
	DstId uint
}

type AddNewMember struct {
	PageArg
	//MemberId uint
	DstName string
}
