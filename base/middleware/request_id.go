package middleware

import (
	"qqlx/base/constant"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set(constant.TraceID, requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

func GetTraceID(c *gin.Context) string {
	if requestID, exists := c.Get(constant.TraceID); exists {
		return requestID.(string)
	}
	return ""
}
