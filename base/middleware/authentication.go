package middleware

import (
	"errors"
	"qqlx/base/apierrs"
	"qqlx/base/constant"
	"qqlx/base/handler"
	"qqlx/pkg/jwt"

	"strings"

	"github.com/gin-gonic/gin"
)

var ErrHeaderEmpty = errors.New("auth in the request header is empty")
var ErrHeaderMalformed = errors.New("the auth format in the request header is incorrect")

// Authentication 基于JWT的认证中间件
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			handler.ResponseForbidden(c, apierrs.NewAuthError(ErrHeaderEmpty))
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			handler.ResponseForbidden(c, apierrs.NewAuthError(ErrHeaderMalformed))
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			handler.ResponseForbidden(c, apierrs.NewAuthError(err))
			return
		}
		c.Set(constant.AuthMidwareKey, mc)
		c.Next()
	}
}
