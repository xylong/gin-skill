package dto

type (
	RegisterReq struct {
		Phone    string `form:"phone" json:"phone" binding:"required"`
		Name     string `form:"user" json:"user" binding:"required"`
		Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
	}

	LoginReq struct {
		Phone    string `form:"user" json:"user" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
)
