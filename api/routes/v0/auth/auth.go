// Copyright © 2020. Drew Lee. All rights reserved.

package auth

import (
	"KarlMalone/pkg/app"
	"KarlMalone/pkg/e"
	"KarlMalone/pkg/logger"
	"KarlMalone/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Get token
// @Description Get a token by member_id in headers
// @Tags auth
// @Param Authorization header string true "Basic [base64-MemberID]"
// @Produce  json
// @Success 201 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 401 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v0/auth/token [get]
func GetToken(c *gin.Context) {
	r := app.Gin{C: c}
	memberId := c.GetInt64("member_id")
	if memberId < 0 {
		r.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	// ??? 去数据库找一下是否存在该memberId? 大量请求会不会导致数据库io性能问题?
	// 这里暂定不去查数据库了，过来一个memberId，就给他一个token
	token, err := util.GenerateToken(memberId)
	if err != nil {
		logger.Error("GetToken error", zap.Int64("member_id", memberId), zap.Error(err))
		r.Response(http.StatusBadRequest, e.ERROR_AUTH_GET_TOKEN_FAIL, nil)
		return
	}

	r.Response(http.StatusCreated, e.SUCCESS, gin.H{"token": token})
}

// @Summary Parse token
// @Description Parse token to member_id
// @Tags auth
// @Produce  json
// @Param token query string true "Token"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 401 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v0/auth/parse_token [get]
func ParseToken(c *gin.Context) {
	r := app.Gin{C: c}
	token := c.Query("token")
	if token == "" {
		r.Response(http.StatusBadRequest, e.ERROR_AUTH_NO_TOKEN, nil)
		return
	}

	claims, err := util.ParseToken(token)
	if err != nil {
		logger.Error("ParseToken error", zap.String("token", token), zap.Error(err))
		r.Response(http.StatusBadRequest, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	r.Response(http.StatusOK, e.SUCCESS, gin.H{"member_id": claims.MemberId})
}
