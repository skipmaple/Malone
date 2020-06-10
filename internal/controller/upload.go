// Copyright © 2020. Drew Lee. All rights reserved.

package controller

import (
	"KarlMalone/pkg/logger"
	"KarlMalone/pkg/util"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
)

func init() {
	if err := os.MkdirAll("./resources", os.ModePerm); err != nil {
		logger.Error("upload controller mkdir error", zap.String("reason", err.Error()))
	}
}

func FileUpload(w http.ResponseWriter, r *http.Request) {
	UploadLocal(w, r)
}

func UploadLocal(w http.ResponseWriter, r *http.Request) {
	srcFile, head, err := r.FormFile("file")
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}

	suffix := ".png"
	srcFilename := head.Filename
	splitMsg := strings.Split(srcFilename, ".")
	if len(splitMsg) > 1 {
		suffix = "." + splitMsg[len(splitMsg)-1]
	}
	fileType := r.FormValue("filetype")
	if len(fileType) > 0 {
		suffix = fileType
	}
	filename := fmt.Sprintf("%d%s%s", time.Now().Unix(), util.GenRandomStr(32), suffix)

	filepath := "./resources/" + filename
	dstFile, err := os.Create(filepath)
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}

	util.RespOk(w, filepath, "")
}
