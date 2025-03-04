package handler

import (
	"net/http"
	"qqlx/base/apierrs"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, newResponse(WithData(data)))
}

func ResponseNotFound(c *gin.Context, err error) {
	c.Set("error", err)
	c.JSON(http.StatusNotFound, conversionError(err))
}

func ResponseServerError(c *gin.Context, err error) {
	c.Set("error", err)
	c.JSON(http.StatusInternalServerError, conversionError(err))
}

func ResponseUnauthorized(c *gin.Context, err error) {
	c.Set("error", err)
	c.JSON(http.StatusUnauthorized, conversionError(err))
}

func ResponseForbidden(c *gin.Context, err error) {
	c.Abort()
	c.Set("error", err)
	c.JSON(http.StatusForbidden, conversionError(err))
}

func ResponParamsError(c *gin.Context, err error) {
	c.Set("error", err)
	c.JSON(http.StatusBadRequest, conversionError(err))
}

// BindAndCheck 绑定参数
func BindAndCheck(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBind(req); err != nil {
		apierr := apierrs.NewParamsError(err)
		ResponParamsError(c, apierr)
		return true
	}
	return false
}

// conversionError 转换错误后设置响应体
func conversionError(err error) *resp {
	apierr, ok := err.(*apierrs.ApiError)
	if !ok {
		return newResponseForErr(WithCode(apierrs.ErrNotApiErr), WithErr(err.Error()))
	}
	return newResponseForErr(WithCode(apierr.Code), WithMessage(apierr.Msg), WithErr(apierr.Err.Error()))
}
