package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 加密
func Md5(str []byte) string {
	m := md5.New()
	m.Write(str)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
