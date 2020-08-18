// Copyright © 2020. Drew Lee. All rights reserved.

package v0

import (
	"KarlMalone/api/routes/v0/auth"
	"KarlMalone/api/routes/v0/members"

	"github.com/gin-gonic/gin"
)

func Handler(r *gin.Engine) {
	v0 := r.Group("/v0")
	{
		// 登录 和 注册
		members.Handler(v0)

		auth.Handler(v0)
	}
}
