// Copyright © 2020. Drew Lee. All rights reserved.

package auth

import (
	"KarlMalone/middleware"

	"github.com/gin-gonic/gin"
)

func Handler(r *gin.RouterGroup) {
	{
		r.GET("/auth/parse_token", ParseToken)
		// 使用该中间件之后的所有路由 请求头里需要包含 member_id 才能访问
		r.Use(middleware.SetMemberIdInContext())
		r.GET("/auth/token", GetToken)
		// 这个路由返回用户信息（包含主题，语言等)
		r.GET("/auth/member_app_info", MemberAppInfo)
	}
}
