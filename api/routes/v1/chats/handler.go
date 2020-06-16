// Copyright © 2020. Drew Lee. All rights reserved.

package chats

import "github.com/gin-gonic/gin"

func Handler(r *gin.RouterGroup) {
	{
		r.POST("/chats/chat", chat)
	}
}
