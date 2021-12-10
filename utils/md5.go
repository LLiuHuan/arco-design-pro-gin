package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5V 简单md5加密
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string
func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
