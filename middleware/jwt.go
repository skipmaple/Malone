// Copyright © 2020. Drew Lee. All rights reserved.

package middleware

import (
	"KarlMalone/pkg/app"
	"KarlMalone/pkg/e"
	"KarlMalone/pkg/util"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := app.Gin{C: c}
		var claims util.Claims
		var code int
		var data interface{}

		code = e.SUCCESS
		token := strings.TrimSpace(c.GetHeader("Authorization"))
		if token == "" {
			code = e.ERROR_AUTH_NO_TOKEN
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			r.Response(http.StatusUnauthorized, code, data)

			c.Abort()
			return
		}

		setMemberIdInContext(c, claims)
		c.Next()
	}
}

func setMemberIdInContext(c *gin.Context, claims util.Claims) {
	c.Set("member_id", claims.MemberId)
}
