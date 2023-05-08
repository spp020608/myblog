package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"myblog/pkg/config"

	"time"
)

// CustomSecret 用于加盐的字符串
var customSecret = []byte("Yusheng")

// TokenExpireDuration 过期时间
const TokenExpireDuration = time.Hour * 24

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的 jwt.RegisteredClaims 只包含了官方字段
// 假设我们这里需要额外记录一个 username 字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

//
// GenerateToken 生成 token 字符串
//	@parameter	username 登录的时候适用的用户名
//	@return		string 生成的 token 字符串
//	@return		error 有错误发生的时候返回的错误信息
//
func GenerateToken(username string) (string, error) {
	cfg := config.Get().JwtConfig
	// 从配置文件中获取过期时间(这个是一个时长, 单位为秒)
	expire := cfg.Expire
	// 计算持续时间段
	duration := time.Second * time.Duration(expire)

	// 创建一个我们自己的声明
	claims := CustomClaims{
		username, // 自定义字段
		jwt.RegisteredClaims{
			// 配置文件中存放的过期时间是秒为单位, 所以在这里需要 time.Second * time.Duration(expire)
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)), // 从配置文件中读取 Jwt 过期参数
			Issuer:    "YuSheng",                          // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)


/*	cfg.Secret 是从配置文件中读取的一个字符串，
	但是在 GenerateToken 函数中它被用来作为 JWT 的签名秘钥。
	JWT 签名秘钥必须是字节数组类型，
	因此需要将字符串类型的 cfg.Secret 转换为字节数组类型*/
	secret := []byte(cfg.Secret)


	// 使用指定的 secret 签名并获得完整的编码后的字符串 token
	return token.SignedString(secret) // 使用配置文件中的参数进行签名

}

//
// ParseJwtToken 解析 Token 字符串
//	@parameter	tokenString 要解析的 Token 字符串
//	@return		*CustomClaims 返回自定义的 Claims 对象
//	@return		error 如果有错误发生
//
func ParseJwtToken(tokenString string) (*CustomClaims, error) {
	// 解析 token
	// 如果是自定义 Claim 结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			// 直接使用标准的 Claim 则可以直接使用 Parse 方法
			// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
			secret := config.Get().JwtConfig.Secret // 从配置文件中获取加密盐
			return secret, nil
		},
	)
	if err != nil {
		return nil, err
	}
	// 对 token 对象中的 Claim 进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
