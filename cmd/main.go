// Copyright © 2020. Drew Lee. All rights reserved.

package main

import (
	v1 "KarlMalone/api/routes/v1"
	"KarlMalone/cmd/base"
	"KarlMalone/pkg/logger"

	_ "KarlMalone/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Malone API
// @version 1.0-Beta
// @contact.name Drew Lee
// @contact.email skipmaple@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	logger.Info("start ...")

	base.Server(func(r *gin.Engine) {
		v1.Handler(r)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	})

	logger.Info("route start success.")
}
