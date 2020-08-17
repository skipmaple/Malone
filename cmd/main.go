// Copyright © 2020. Drew Lee. All rights reserved.

package main

import (
	v0 "KarlMalone/api/routes/v0"
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
// @license.name MIT License
// @license.url https://github.com/UncleMaple/Malone/blob/master/LICENSE.md

func main() {
	logger.Info("start ...")

	base.Server(func(r *gin.Engine) {
		v0.Handler(r)
		v1.Handler(r)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	})

	logger.Info("route start success.")
}
