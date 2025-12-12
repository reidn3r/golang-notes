package singleton

var instance *SingleUser

type SingleUser struct {
	name string
	age  uint
}

type SingletonInterface interface {
	SetName(name string) *SingleUser
	SetAge(age uint) *SingleUser
}

func GetObject() *SingleUser {
	if instance == nil {
		instance = &SingleUser{}
	}
	return instance
}

func (u *SingleUser) SetName(name string) *SingleUser {
	u.name = name
	return u
}

func (u *SingleUser) SetAge(age uint) *SingleUser {
	u.age = age
	return u
}
