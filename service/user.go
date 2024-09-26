package service

import (
	"gin-skill/dao"
)

func GetUserById(id int64) (any, error) {
	u := dao.User
	return u.Where(u.ID.Eq(id)).Preload(dao.User.Profile).First()
}
