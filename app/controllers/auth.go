package controllers

import (
	"gin-skill/dto"
	"github.com/gin-gonic/gin"
)

// Register 注册
func Register(ctx *gin.Context) (any, error) {
	var (
		err error
		req dto.RegisterReq
	)

	if err = ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	//t := time.Date(2000, 2, 2, 0, 0, 0, 0, time.Local)
	//users := []*models.User{
	//	{
	//		Phone:    "13512341234",
	//		Email:    "summer@gmail.com",
	//		Name:     "summer",
	//		Nickname: "夏天",
	//		Profile: models.UserProfile{
	//			Gender:   1,
	//			Level:    1,
	//			Birthday: &t,
	//		},
	//		Address: []models.Address{
	//			{Province: "", City: "上海", Address: "汤臣一品"},
	//			{Province: "", City: "北京", District: "朝阳区", Address: "朝阳区"},
	//		},
	//	},
	//}
	//
	//err := dao.User.Create(users...)
	//if err != nil {
	//	return nil, err
	//}

	return gin.H{"a": "app"}, nil
}
