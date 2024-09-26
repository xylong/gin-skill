package user

import "time"

type (
	Attr  func(user *User)
	Attrs []Attr
)

func (a Attrs) Apply(user *User) {
	for _, f := range a {
		f(user)
	}
}

func WithId(id int64) Attr {
	return func(user *User) {
		user.ID = id
	}
}

func WithAvatar(avatar string) Attr {
	return func(user *User) {
		user.Avatar = avatar
	}
}

func WithName(name string) Attr {
	return func(user *User) {
		user.Name = name
	}
}

func WithGender(gender int) Attr {
	return func(user *User) {
		user.Gender = gender
	}
}

func WithBirthDay(birthday string) Attr {
	return func(user *User) {
		t, err := time.Parse(time.DateOnly, birthday)
		if err != nil {
			return
		}

		user.BirthDay = t
	}
}
