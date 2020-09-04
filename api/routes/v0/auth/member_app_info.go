// Copyright © 2020. Drew Lee. All rights reserved.

package auth

import (
	"KarlMalone/internal/service"
	"KarlMalone/pkg/app"
	"KarlMalone/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Member app info
// @Description Get token by member_id
// @Tags auth
// @Param member_id header string true "Member ID"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 401 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v0/auth/member_app_info [get]
func MemberAppInfo(c *gin.Context) {
	r := app.Gin{C: c}
	memberId := c.GetInt64("member_id")
	if memberId <= 0 {
		r.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	// 暂时先返回member信息，后续补全 语言，主题，设备信息等
	m := service.Member{ID: memberId}
	member, err := m.Find()
	if err != nil {
		r.Response(http.StatusNotFound, e.ERROR_NOT_EXIST_MEMBER, nil)
		return
	}

	r.Response(http.StatusOK, e.SUCCESS, member)
}
