package common

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func Base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}
func Base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

const (
	md5Table = "!@#$%^&*()_+-=123456789"
)

func MD5EncodeString(src string) string {
	return hex.EncodeToString(MD5EncodeBytes([]byte(src)))
}

func MD5EncodeBytes(src []byte) []byte {
	h := md5.New()
	h.Write(src)

	cipherStr := h.Sum([]byte(md5Table))

	return cipherStr
}
