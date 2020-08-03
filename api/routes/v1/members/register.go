// Copyright © 2020. Drew Lee. All rights reserved.

package members

import (
	"KarlMalone/internal/service"
	"KarlMalone/pkg/app"
	"KarlMalone/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary register
// @Description member register
// @Accept multipart/form-data
// @Produce  json
// @Param phone_num formData string true "PhoneNum"
// @Param email formData string false "Email"
// @Param password formData string true "Password"
// @Param nickname formData string true "Nickname"
// @Param gender formData string true "Gender(Male Female Unknown)" Enums(M, F, U) Default(U)
// @Param avatar formData string false "Avatar"
// @Param memo formData string false "Memo"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/members/register [post]
func register(c *gin.Context) {
	r := app.Gin{C: c}

	m := service.Member{}
	if err := c.ShouldBind(&m); err != nil {
		r.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}

	member, err := m.Register()
	if err != nil {
		r.Response(http.StatusInternalServerError, e.ERROR_REGISTER_MEMBER, nil)
	} else {
		r.Response(http.StatusOK, e.SUCCESS, member)
	}
}
