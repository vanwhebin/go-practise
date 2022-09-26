package model

import (
	"crypto/md5"
	"encoding/hex"
)

// GeneratePasswordHash : Use MD5
func GeneratePasswordHash(pwd string) string {
	return Md5(pwd)
}

func Md5(origin string) string {
	hasher := md5.New()
	hasher.Write([]byte(origin))
	pwdHash := hex.EncodeToString(hasher.Sum(nil))
	return pwdHash
}
