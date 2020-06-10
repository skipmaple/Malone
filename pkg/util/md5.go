// Copyright © 2020. Drew Lee. All rights reserved.

package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

func Md5EncodeUpper(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

func MakePwd(plainPwd, salt string) string {
	return Md5Encode(plainPwd + salt)
}

func ValidatePwd(plainPwd, salt, pwd string) bool {
	return Md5Encode(plainPwd+salt) == pwd
}
