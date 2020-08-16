// Copyright © 2020. Drew Lee. All rights reserved.

package members

import "github.com/gin-gonic/gin"

func Handler(r *gin.RouterGroup) {
	{
		r.POST("/members/login", login)
		r.POST("/members/register", register)
		r.GET("/members/find", find)
	}
}
