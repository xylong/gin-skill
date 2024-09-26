package dto

type (
	UserInfoReq struct {
		Id int64 `uri:"id" binding:"required,gt=0"`
	}
)

type (
	SimpleUser struct {
		Id        int64  `json:"id"`
		Name      string `json:"name"`
		Nickname  string `json:"nickname"`
		Avatar    string `json:"avatar"`
		Level     int64  `json:"level"`
		Signature string `json:"signature"`
		CreatedAt string `json:"created_at"`
	}
)
