package services

import (
	"fmt"
	"gin-skill/app/dao"
	"gin-skill/app/models"
	"gin-skill/dto"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"time"
)

type userService struct{}

var (
	UserService = new(userService)
)

// Me 获取个人信息
func (s *userService) Me(id string) (error, *dto.SimpleUser) {
	var (
		err        error
		user       *models.User
		simpleUser dto.SimpleUser
	)

	u := dao.User
	if user, err = u.Where(u.ID.Eq(cast.ToInt64(id))).Preload(dao.User.Profile).First(); err != nil {
		return fmt.Errorf("用户不存在"), nil
	}

	_ = copier.Copy(&simpleUser, &user)
	simpleUser.Level = user.Profile.Level
	simpleUser.Signature = user.Profile.Signature
	simpleUser.CreatedAt = user.CreatedAt.Format(time.DateTime)

	return nil, &simpleUser
}
