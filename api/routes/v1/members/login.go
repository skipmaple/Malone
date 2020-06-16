// Copyright © 2020. Drew Lee. All rights reserved.

package members

import (
	"KarlMalone/internal/service"
	"KarlMalone/pkg/app"
	"KarlMalone/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

// member login
func login(c *gin.Context) {
	r := app.Gin{C: c}

	phoneNum := c.PostForm("phone_num")
	plainPwd := c.PostForm("password")

	// validate parameters
	if len(phoneNum) == 0 || len(plainPwd) == 0 {
		r.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
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
