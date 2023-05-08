package util

import "golang.org/x/crypto/bcrypt"

//
// GenerateFromPassword 根据原始密码字符串生成加密密码字符串
//	@parameter	password 原始密码字符串
//	@return		string 加密之后的密码
//	@return		error 错误信息
//
func GenerateFromPassword(password string) (string, error) {
	// 第二个参数是进行哈希的次数，这里采用了默认值 10, 数字越大生成的密码速度越慢，成本越大。但是更安全
	// bcrypt 每次生成的编码是不同的，较于 md5 更安全
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

//
// CompareHashAndPassword 比较加密密码和名文密码
//	@parameter	hashedPassword 已经加密过的密码字符串
//	@parameter	password 原始密码字符串
//	@return		error
//
func CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
