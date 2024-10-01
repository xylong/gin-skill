package request

import (
	"gin-skill/app/common/response"
	"github.com/gin-gonic/gin"
)

// Validator 验证函数
func Validate(ctx *gin.Context, validator Validator) {
	if err := ctx.ShouldBind(validator); err != nil {
		response.ValidateFail(ctx, GetErrorMsg(validator.GetMessages(), err))
	}
}
