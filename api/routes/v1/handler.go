// Copyright © 2020. Drew Lee. All rights reserved.

package v1

import (
	"KarlMalone/api/routes/v1/attaches"
	"KarlMalone/api/routes/v1/chats"
	"KarlMalone/api/routes/v1/contacts"
	"KarlMalone/api/routes/v1/members"
	"KarlMalone/middleware"

	"github.com/gin-gonic/gin"
)

func Handler(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.Use(middleware.JWT())
	{
		attaches.Handler(v1)
		contacts.Handler(v1)
		chats.Handler(v1)
		members.Handler(v1)
	}
}
