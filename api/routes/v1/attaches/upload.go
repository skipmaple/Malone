// Copyright © 2020. Drew Lee. All rights reserved.

package attaches

import (
	"KarlMalone/pkg/app"
	"KarlMalone/pkg/e"
	"KarlMalone/pkg/logger"
	"KarlMalone/pkg/util"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func init() {
	if err := os.MkdirAll("./resources", os.ModePerm); err != nil {
		logger.Error("upload controller mkdir error", zap.String("reason", err.Error()))
	}
}

func fileUpload(c *gin.Context) {
	r := app.Gin{C: c}
	file, err := c.FormFile("file")
	if err != nil {
		r.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	suffix := ".png"
	srcFilename := file.Filename
	splitMsg := strings.Split(srcFilename, ".")
	if len(splitMsg) > 1 {
		suffix = "." + splitMsg[len(splitMsg)-1]
	}

	fileType := c.PostForm("file-type")
	if len(fileType) > 0 {
		suffix = fileType
	}
	filename := fmt.Sprintf("%d%s%s", time.Now().Unix(), util.GenRandomStr(32), suffix)
	filepath := "./resources/" + filename

	err = c.SaveUploadedFile(file, filepath)
	if err != nil {
		r.Response(http.StatusInternalServerError, e.ERROR, nil)
	} else {
		r.Response(http.StatusOK, e.SUCCESS, nil)
	}
}
