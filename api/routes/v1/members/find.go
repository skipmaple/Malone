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

// @Summary Find member by member_id
// @Description Find member by member_id
// @Tags member
// @Produce  json
// @Param member_id query string true "MemberId"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/members/find [get]
func find(c *gin.Context) {
	r := app.Gin{C: c}
	memberId, _ := strconv.ParseInt(c.Query("member_id"), 10, 64)
	m := service.Member{
		ID: memberId,
	}

	member := m.Find()
	r.Response(http.StatusOK, e.SUCCESS, member)
}
