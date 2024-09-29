package services

import (
	"fmt"
	"gin-skill/app/dao"
	"gin-skill/app/models"
	"gin-skill/app/requests"
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
func (s *authService) Register(req request.RegisterReq) (error, *models.User) {
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

// Login 登录
func (s *authService) Login(req request.LoginReq) (error, *models.User) {
	var (
		err  error
		u    = dao.User
		user *models.User
	)

	user, err = u.Where(u.Phone.Eq(req.Phone)).Preload(u.Profile).First()
	if err != nil || !utils.BcryptMakeCheck([]byte(req.Password), user.Profile.Password) {
		return fmt.Errorf("用户不存在或密码错误"), nil
	}

	return nil, user
}
