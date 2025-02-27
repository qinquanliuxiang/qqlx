package jwt

import (
	"fmt"
	"time"
	"qqlx/base/apierrs"
	"qqlx/base/conf"
	"qqlx/base/constant"
	"qqlx/base/reason"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	UserID   uint   `json:"userId,omitempty"`
	UserName string `json:"userName,omitempty"`
	*jwt.RegisteredClaims
}

// NewClaims creates a new instance of MyCustomClaims with the given userID and zhName.
func NewClaims(userID uint, userName string) *MyCustomClaims {
	now := time.Now()
	return &MyCustomClaims{
		UserID:   userID,
		UserName: userName,
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    conf.GetJwtIssuer(),
			ExpiresAt: jwt.NewNumericDate(now.Add(conf.GetJwtExpirationTime())),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
}

func (c *MyCustomClaims) GenerateToken() (token string, err error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err = claims.SignedString([]byte(conf.GetJwtSecret()))
	if err != nil {
		return "", apierrs.NewGenerateTokenError(fmt.Errorf("failed to generate token, %w", err))
	}
	return token, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	var myCustomClaims MyCustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &myCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.GetJwtSecret()), nil
	})
	if err != nil {
		return nil, apierrs.NewParseTokenError(fmt.Errorf("failed to parse token, %w", err))
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok {
		return claims, nil
	}
	return nil, apierrs.NewParseTokenError(reason.ErrTokenMode)
}

// GetMyCustomClaims 从gin.Context获取MyCustomClaims
func GetMyCustomClaims(c *gin.Context) (*MyCustomClaims, error) {
	cl, ok := c.Get(constant.AuthMidwareKey)
	if !ok {
		return nil, apierrs.NewAuthError(reason.ErrHeaderEmpty)
	}
	myCustomClaims, ok := cl.(*MyCustomClaims)
	if !ok {
		return nil, apierrs.NewAuthError(reason.ErrTokenMode)
	}
	return myCustomClaims, nil
}
