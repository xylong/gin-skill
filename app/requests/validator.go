package request

import (
	"github.com/go-playground/validator/v10"
)

type (
	Validator interface {
		GetMessages() ValidatorMessages
	}

	// ValidatorMessages 验证消息
	ValidatorMessages map[string]string
)

// GetErrorMsg 获取错误信息
func GetErrorMsg(request interface{}, err error) string {
	_, isValidator := request.(Validator)

	if _, ok := err.(validator.ValidationErrors); ok && isValidator {
		for _, v := range err.(validator.ValidationErrors) {
			if message, exist := request.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exist {
				return message
			}

			return v.Error()
		}
	}

	return "参数错误"
}
