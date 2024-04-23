package hashed

import (
	"crypto/md5"
	"encoding/hex"
)


func HashPassword(password string) string{
	hash := md5.Sum([]byte(password))
	hashedPass := hex.EncodeToString(hash[:])
	return hashedPass
}