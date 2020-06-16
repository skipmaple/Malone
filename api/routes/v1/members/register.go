// Copyright © 2020. Drew Lee. All rights reserved.

package members

import (
	"KarlMalone/internal/service"
	"KarlMalone/pkg/app"
	"KarlMalone/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

// member register
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
