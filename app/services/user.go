package services

import (
	"fmt"
	"gin-skill/app/dao"
	"gin-skill/app/models"
	"gin-skill/dto"
	"github.com/jinzhu/copier"
	"time"
)

const (
	UserNotFound = 10001
)

func GetSimpleUser(id int64) (any, error) {
	var (
		err        error
		user       *models.User
		simpleUser dto.SimpleUser
	)

	u := dao.User
	if user, err = u.Where(u.ID.Eq(id)).Preload(dao.User.Profile).First(); err != nil {
		return nil, fmt.Errorf("用户不存在")
	}

	_ = copier.Copy(&simpleUser, &user)
	simpleUser.Level = user.Profile.Level
	simpleUser.Signature = user.Profile.Signature
	simpleUser.CreatedAt = user.CreatedAt.Format(time.DateTime)

	return &simpleUser, nil
}
