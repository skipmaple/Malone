// Copyright © 2020. Drew Lee. All rights reserved.

package middleware

import (
	"KarlMalone/pkg/app"
	"KarlMalone/pkg/e"
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetMemberIdInContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := app.Gin{C: c}
		authorization := c.GetHeader("Authorization")
		if !strings.HasPrefix(authorization, "Basic ") {
			r.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
			c.Abort()
			return
		}

		basic := strings.TrimPrefix(authorization, "Basic ")
		memberId, err := base64.StdEncoding.DecodeString(basic)
		if err != nil || string(memberId) == "" {
			r.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
			c.Abort()
			return
		}

		mId, err := strconv.ParseInt(string(memberId), 10, 64)
		if err != nil {
			r.Response(http.StatusBadRequest, e.ERROR_AHTH_INVALID_HEADERS, nil)
			c.Abort()
			return
		}
		c.Set("member_id", mId)
		c.Next()
	}
}
