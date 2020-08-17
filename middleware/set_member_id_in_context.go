// Copyright © 2020. Drew Lee. All rights reserved.

package middleware

import (
	"KarlMalone/pkg/app"
	"KarlMalone/pkg/e"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetMemberIdInContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := app.Gin{C: c}
		// !!! 优化可以考虑传输加密的member_id，在服务端解析一下
		memberId := c.GetHeader("member_id")
		if memberId == "" {
			r.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
			c.Abort()
			return
		}

		mId, err := strconv.ParseInt(memberId, 10, 64)
		if err != nil {
			r.Response(http.StatusBadRequest, e.ERROR_AHTH_INVALID_HEADERS, nil)
			c.Abort()
			return
		}
		c.Set("member_id", mId)
		c.Next()
	}

}
