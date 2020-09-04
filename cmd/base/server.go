// Copyright © 2020. Drew Lee. All rights reserved.

package base

import (
	"KarlMalone/config"
	"KarlMalone/pkg/logger/ginzap"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server provides base server pre-config
func Server(op func(r *gin.Engine)) {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(ginzap.Ginzap())
	r.Use(ginzap.RecoveryWithZap(true))
	op(r)

	srv := &http.Server{
		Addr:    config.Server.Port,
		Handler: r,
	}

	log.Fatal(srv.ListenAndServe())
}
