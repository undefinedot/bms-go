package utils

import "golang.org/x/crypto/bcrypt"

// BcryptHash 加密密码
func BcryptHash(pwd string) string {
	password, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(password)
}

// CheckPassword 校验密码
func CheckPassword(pwd, target string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(target))
	return err == nil
}
