package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// BcryptMake 密码加密
func BcryptMake(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

// BcryptMakeCheck 密码校验
func BcryptMakeCheck(pwd []byte, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	return bcrypt.CompareHashAndPassword(byteHash, pwd) == nil
}
