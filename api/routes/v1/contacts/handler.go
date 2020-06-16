// Copyright © 2020. Drew Lee. All rights reserved.

package contacts

import "github.com/gin-gonic/gin"

func Handler(r *gin.RouterGroup) {
	{
		r.POST("/contacts/add_friend", addFriend)
		r.POST("/contacts/create_group", createGroup)
		r.POST("/contacts/join_group", joinGroup)
		r.GET("/contacts/load_friend", loadFriend)
		r.GET("/contacts/load_group", loadGroup)
	}
}
