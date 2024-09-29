package services

import (
	"fmt"
	"gin-skill/app/common/request"
	"gin-skill/app/dao"
	"gin-skill/app/models"
	"gin-skill/utils"
	"go.uber.org/zap"
)

type authService struct {
}

var (
	// AuthService 授权
	AuthService = new(authService)
)

// Register 注册
func (userService *authService) Register(req request.RegisterReq) (error, *models.User) {
	var (
		err  error
		u    = dao.User
		user *models.User
	)

	if user, err = u.Where(u.Phone.Eq(req.Phone)).Select(u.ID).First(); user != nil {
		return fmt.Errorf("手机号已存在"), nil
	}

	user = &models.User{
		Phone: req.Phone,
		Name:  req.Name,
	}
	user.Profile.Password = utils.BcryptMake([]byte(req.Password))
	if err = u.Create(user); err != nil {
		zap.L().Error("register", zap.Error(err))
		return fmt.Errorf("注册失败"), nil
	}

	return nil, user
}
