package lib

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	if str == "" {
		return ""
	}
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
