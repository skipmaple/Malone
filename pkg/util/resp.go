// Copyright © 2020. Drew Lee. All rights reserved.

package util

import (
	"KarlMalone/pkg/logger"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type H struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Rows  interface{} `json:"rows,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}

func RespOk(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, msg)
}

func RespOkList(w http.ResponseWriter, lists interface{}, total interface{}) {
	RespList(w, 0, lists, total)
}

func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	res, err := json.Marshal(h)
	if err != nil {
		logger.Panic("respList util json marshal error", zap.String("reason", err.Error()))
	}

	_, _ = w.Write(res)
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	res, err := json.Marshal(resp)
	if err != nil {
		logger.Panic("resp util json marshal error", zap.String("reason", err.Error()))
	}

	_, _ = w.Write(res)
}
