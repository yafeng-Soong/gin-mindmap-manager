package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const (
	TIMES = 2
)

func MD5V(str []byte, salt []byte, iteration int) string {
	h := md5.New()
	h.Write(salt)
	h.Write(str)
	str = h.Sum(nil)
	for i := 0; i < iteration-1; i++ {
		h.Reset()
		h.Write(str)
		str = h.Sum(nil)
	}
	return hex.EncodeToString(str)
}
