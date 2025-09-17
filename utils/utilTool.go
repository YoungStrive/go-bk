package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 将str md5加密
func Md5Str(str string) string {
	data := []byte(str)
	sum := md5.Sum(data)
	md5Str := hex.EncodeToString(sum[:])
	return md5Str

}
