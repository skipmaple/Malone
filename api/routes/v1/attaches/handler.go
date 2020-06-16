// Copyright © 2020. Drew Lee. All rights reserved.

package attaches

import "github.com/gin-gonic/gin"

func Handler(r *gin.RouterGroup) {
	{
		r.POST("/attaches/upload", fileUpload)
	}
}
