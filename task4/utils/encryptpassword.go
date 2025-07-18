package utils

import "golang.org/x/crypto/bcrypt"

//密码加密
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // bcrypt.DefaultCost 是默认的加密成本
	return string(hashedPassword), err
}

//密码核对
func CheckPassword(hashpassword, password string) bool {
	// 比较密码哈希和用户输入的密码
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	return err == nil // 如果错误为nil，则密码匹配
}
