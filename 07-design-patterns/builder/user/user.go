package user

type User struct {
	name   string
	age    uint
	height float64
}

type UserBuilder struct {
	name   string
	age    uint
	height float64
}

type UserBuilderInterface interface {
	SetName(name string) UserBuilderInterface
	SetAge(age uint) UserBuilderInterface
	SetHeight(h float64) UserBuilderInterface
	Build() *User
}

func CreateBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (u *UserBuilder) SetName(name string) UserBuilderInterface {
	u.name = name
	return u
}

func (u *UserBuilder) SetAge(age uint) UserBuilderInterface {
	u.age = age
	return u
}

func (u *UserBuilder) SetHeight(h float64) UserBuilderInterface {
	u.height = h
	return u
}

func (u *UserBuilder) Build() *User {
	return &User{
		name:   u.name,
		age:    u.age,
		height: u.height,
	}
}
