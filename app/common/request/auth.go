package request

// RegisterReq 注册
type RegisterReq struct {
	Phone    string `form:"phone" json:"phone" binding:"required,mobile"`
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
}

// GetMessages 自定义错误信息
func (register RegisterReq) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":     "用户名称不能为空",
		"phone.required":    "手机号码不能为空",
		"phone.mobile":      "手机号码异常",
		"password.required": "用户密码不能为空",
		"password.min":      "用户密码最小6位",
		"password.max":      "用户密码最大20位",
	}
}

// LoginReq 登录
type LoginReq struct {
	Phone    string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
