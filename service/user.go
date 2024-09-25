package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(id int64) (gin.H, error) {
	if id > 10 {
		return gin.H{"msg": "ok"}, nil
	}

	return nil, fmt.Errorf("test error")
}
