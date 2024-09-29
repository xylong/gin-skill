package global

// CustomError 自定义错误
type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

// CustomErrors 自定义错误
type CustomErrors struct {
	BusinessError CustomError
	ValidateError CustomError
	TokenError    CustomError
}

var Errors = CustomErrors{
	BusinessError: CustomError{40000, "业务错误"},
	ValidateError: CustomError{42200, "请求参数错误"},
	TokenError:    CustomError{40100, "登录授权失效"},
}
