package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func MakePassword(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateSession(user_id int64) string {
	h := md5.New()
	h.Write([]byte(strconv.FormatInt(user_id, 10)))
	return hex.EncodeToString(h.Sum(nil))
}
