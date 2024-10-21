package core

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

func formatClaims(claims jwt.MapClaims) string {
	var formattedClaims strings.Builder
	for key, value := range claims {
		switch v := value.(type) {
		case float64:
			// 将 float64 转换为整数格式以避免科学计数法
			formattedClaims.WriteString(fmt.Sprintf("%s: %.0f\n", key, v))
		default:
			// 对于其他类型的值，直接输出
			formattedClaims.WriteString(fmt.Sprintf("%s: %v\n", key, v))
		}
	}
	return formattedClaims.String()
}
