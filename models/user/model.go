package user

import "time"

type User struct {
	ID       int64     `json:"id"`
	Avatar   string    `json:"avatar"`
	Name     string    `json:"name"`
	Gender   int       `json:"gender"`
	BirthDay time.Time `json:"birthday,omitempty"`
}

func NewUser(attrs ...Attr) *User {
	u := &User{}
	Attrs(attrs).Apply(u)
	return u
}

// Mutate 设置属性
func (u *User) Mutate(attrs ...Attr) *User {
	Attrs(attrs).Apply(u)
	return u
}
