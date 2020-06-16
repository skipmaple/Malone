// Copyright © 2020. Drew Lee. All rights reserved.

package main

import (
	v1 "KarlMalone/api/routes/v1"
	"KarlMalone/cmd/base"
	"KarlMalone/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("start ...")

	base.Server(func(r *gin.Engine) {
		v1.Handler(r)
	})

	logger.Info("route start success.")
}
