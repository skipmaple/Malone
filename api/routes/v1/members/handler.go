// Copyright © 2020. Drew Lee. All rights reserved.

package members

import (
	"github.com/gin-gonic/gin"
)

func Handler(r *gin.RouterGroup) {
	{
		r.GET("/members/find", find)
		r.GET("/members/logout", logout)
	}
}
