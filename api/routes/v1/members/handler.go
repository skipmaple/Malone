// Copyright © 2020. Drew Lee. All rights reserved.

package members

import "github.com/gin-gonic/gin"

func Handler(r *gin.RouterGroup) {
	{
		r.GET("/members/login", login)
		r.GET("/members/register", register)
	}
}
