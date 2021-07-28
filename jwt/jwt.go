package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtSecret  = []byte("xxx") // 密钥
	ExpireTime = 30 * time.Second // 令牌有效时间
	Issuer     = "MF"          // 签发人
)

// 生成token的实体类，参数可自定义
type JwtClaims struct {
	UserId   interface{} `json:"user_id"`
	Username interface{} `json:"username"`
	jwt.StandardClaims
}

// 创建token, 参数可自定义
func GenerateToken(userId, username interface{}) (string, error) {
	claims := &JwtClaims{
		UserId:   userId,
		Username: username,
	}
	claims.Issuer = Issuer
	newTime := time.Now()
	claims.IssuedAt = newTime.Unix()
	claims.ExpiresAt = newTime.Add(ExpireTime).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)

	return signedToken, err
}

// 解析token
func ParseToken(strToken string) (*JwtClaims, bool) {
	token, _ := jwt.ParseWithClaims(strToken, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
			return claims, true
		}
	}

	return nil, false
}