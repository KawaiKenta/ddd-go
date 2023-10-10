package entity

import valueobject "github.com/KawaiKenta/ddd-go/value_object"

type User struct {
	Id    int
	Name  *valueobject.UserName
	Email *valueobject.Email
}

func NewUser(name *valueobject.UserName, email *valueobject.Email) *User {
	return &User{
		Id:    0,
		Name:  name,
		Email: email,
	}
}

func (u *User) Equals(other *User) bool {
	return u.Id == other.Id
}

func (u *User) ChangeName(name *valueobject.UserName) {
	u.Name = name
}

func (u *User) ChangeEmail(email *valueobject.Email) {
	u.Email = email
}
