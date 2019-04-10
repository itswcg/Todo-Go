package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MakePassword(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}
