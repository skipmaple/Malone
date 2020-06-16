// Copyright © 2020. Drew Lee. All rights reserved.

package base

import (
	"KarlMalone/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server provides base server pre-config
func Server(op func(r *gin.Engine)) {
	r := gin.Default()
	op(r)

	srv := &http.Server{
		Addr:    config.Server.Port,
		Handler: r,
	}

	log.Fatal(srv.ListenAndServe())
}
