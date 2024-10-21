package core

import (
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/golang-jwt/jwt/v4"
)

func checkWeakKey(jwtToken string, weakKey string) {
	decodedKey, err := base64.StdEncoding.DecodeString(weakKey)

	// 创建一个 JWT 解析器
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// 使用弱密钥进行解密
		return decodedKey, nil
	})

	if err != nil {
		// 检查是否是因为过期引起的错误
		if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors == jwt.ValidationErrorExpired {
			claims, _ := token.Claims.(jwt.MapClaims)
			fmt.Printf("[*] Weak key %s successfully decoded the JWT (expired):\n%v\n", weakKey, formatClaims(claims))
			return
		}
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("[*] Weak key %s successfully decoded the JWT:\n%v\n", weakKey, formatClaims(claims))
	}
}
