// Copyright © 2020. Drew Lee. All rights reserved.

package members

import (
	"KarlMalone/internal/models"
	"KarlMalone/internal/service"
	"KarlMalone/pkg/app"
	"KarlMalone/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary login
// @Description member login
// @Accept multipart/form-data
// @Produce  json
// @Param phone_num formData string true "PhoneNum"
// @Param password formData string true "Password"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/members/login [post]
func login(c *gin.Context) {
	r := app.Gin{C: c}
	member := models.Member{}
	if err := c.ShouldBind(&member); err != nil {
		r.Response(http.StatusInternalServerError, e.ERROR, err)
		return
	}
	phoneNum := member.PhoneNum
	plainPwd := member.Password

	// validate parameters
	if len(phoneNum) == 0 || len(plainPwd) == 0 {
		r.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	m := service.Member{
		PhoneNum: phoneNum,
		Password: plainPwd,
	}
	member, err := m.Login()
	if err != nil {
		r.Response(http.StatusInternalServerError, e.ERROR_LOGIN_MEMBER, nil)
	} else {
		r.Response(http.StatusOK, e.SUCCESS, member)
	}
}
