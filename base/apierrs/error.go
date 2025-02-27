package apierrs

import (
	"fmt"
	"runtime"
)

type ApiError struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg,omitempty"`
	Err   error  `json:"err,omitempty"`
	Stack string `json:"stack"`
}

func (e *ApiError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Msg
}

// Unwrap 实现 Unwrap 方法，允许递归解包底层错误
func (e *ApiError) Unwrap() error {
	return e.Err
}

// Option 定义用于设置 ApiError 的函数类型
type Option func(*ApiError)

// WithMsg 设置错误信息
func WithMsg(msg string) Option {
	return func(e *ApiError) {
		e.Msg = msg
	}
}

// WithErr 设置嵌套错误
func WithErr(err error) Option {
	return func(e *ApiError) {
		e.Err = err
	}
}

// withStack 设置错误堆栈信息
func withStack() Option {
	return func(e *ApiError) {
		_, file, line, _ := runtime.Caller(3)
		e.Stack = fmt.Sprintf("%s:%d", file, line)
	}
}

func WithCode(code int) Option {
	return func(e *ApiError) {
		e.Code = code
	}
}

func WithMsgf(format string, args ...interface{}) Option {
	return func(e *ApiError) {
		e.Msg = fmt.Sprintf(format, args...)
	}
}

// NewError 使用可选参数初始化一个 ApiError
func NewError(code int, options ...Option) *ApiError {
	e := &ApiError{
		Code: code,
	}
	for _, opt := range options {
		opt(e)
	}
	return e
}

func NewParamsError(err error) *ApiError {
	return NewError(
		ErrParamCode,
		WithErr(err),
		withStack(),
	)
}

func NewAuthError(err error) *ApiError {
	return NewError(
		ErrAuthCode,
		WithErr(err),
		withStack(),
	)
}

func NewEncryptError(err error) *ApiError {
	return NewError(
		ErrEncryptCode,
		WithErr(err),
		withStack(),
	)
}

func NewParseTokenError(err error) *ApiError {
	return NewError(
		ErrParseTokenCode,
		WithErr(err),
		withStack(),
	)
}

func NewGenerateTokenError(err error) *ApiError {
	return NewError(
		ErrGenerateTokenCode,
		WithErr(err),
		withStack(),
	)
}

func NewCreateError(err error) *ApiError {
	return NewError(
		ErrCreateCode,
		WithErr(err),
		withStack(),
	)
}

func NewDeleteError(err error) *ApiError {
	return NewError(
		ErrDeleteCode,
		WithErr(err),
		withStack(),
	)
}
func NewListError(err error) *ApiError {
	return NewError(
		ErrListCode,
		WithErr(err),
		withStack(),
	)
}

func NewSaveError(err error) *ApiError {
	return NewError(
		ErrSaveCode,
		WithErr(err),
		withStack(),
	)
}

func NewUpdateError(err error) *ApiError {
	return NewError(
		ErrUpdateCode,
		WithErr(err),
		withStack(),
	)
}

func NewGetError(err error) *ApiError {
	return NewError(
		ErrGetCode,
		WithErr(err),
		withStack(),
	)
}

func NewRedisSetErr(err error) *ApiError {
	return NewError(
		ErrRedisSetCode,
		WithErr(err),
		withStack(),
	)
}

func NewRedisGetErr(err error) *ApiError {
	return NewError(
		ErrRedisGetCode,
		WithErr(err),
		withStack(),
	)
}

func NewRedisDelErr(err error) *ApiError {
	return NewError(
		ErrRedisDelCode,
		WithErr(err),
		withStack(),
	)
}

func NewRedisFlushErr(err error) *ApiError {
	return NewError(
		ErrRedisFlushCode,
		WithErr(err),
		withStack(),
	)
}
