package models

import (
	"crypto/md5"
	"encoding/hex"
)

const passwordLast = "myblog"

func HashPassword(p string) (hp string) {
	h := md5.New()
	h.Write([]byte(p + passwordLast))
	return hex.EncodeToString(h.Sum(nil))
}
