package models

import (
	"github.com/spf13/cast"
)

func (u User) GetUid() string {
	return cast.ToString(u.ID)
}
