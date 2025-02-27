package apierrs

const (
	ErrParamCode = 1000 + iota
	ErrAuthCode
	ErrGenerateTokenCode
	ErrParseTokenCode
	ErrNotApiErr
	ErrEncryptCode
)

const (
	ErrCreateCode = 2000 + iota
	ErrSaveCode
	ErrDeleteCode
	ErrListCode
	ErrUpdateCode
	ErrGetCode
	ErrRedisSetCode
	ErrRedisGetCode
	ErrRedisDelCode
	ErrRedisFlushCode
)
