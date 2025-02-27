package handler

type Option func(*resp)

type resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
	Err  string `json:"err,omitempty"`
}

func WithCode(c int) Option {
	return func(res *resp) {
		res.Code = c
	}
}

func WithMessage(m string) Option {
	return func(res *resp) {
		res.Msg = m
	}
}

func WithData(data any) Option {
	return func(res *resp) {
		res.Data = data
	}
}

func WithErr(err string) Option {
	return func(res *resp) {
		res.Err = err
	}
}

func newResponse(options ...Option) *resp {
	res := &resp{
		Code: 0,
		Msg:  "success",
		Data: "",
	}

	for _, option := range options {
		option(res)
	}
	return res
}

func newResponseForErr(options ...Option) *resp {
	res := &resp{}

	for _, option := range options {
		option(res)
	}
	return res
}
