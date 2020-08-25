// Copyright © 2020. Drew Lee. All rights reserved.

package members

import (
	"KarlMalone/internal/service"
	"KarlMalone/pkg/app"
	"KarlMalone/pkg/e"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Logout member by member_id
// @Description Logout member by member_id
// @Tags member
// @Produce  json
// @Param member_id query string true "MemberId"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/members/logout [get]
func logout(c *gin.Context) {
	r := app.Gin{C: c}
	memberId, _ := strconv.ParseInt(c.Query("member_id"), 10, 64)
	m := service.Member{
		ID: memberId,
	}

	ok := m.Logout()
	if !ok {
		r.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	r.Response(http.StatusOK, e.SUCCESS, nil)
}
