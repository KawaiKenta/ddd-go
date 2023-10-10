package entity

import valueobject "github.com/KawaiKenta/ddd-go/value-object"

type User struct {
	Id    int
	Name  *valueobject.UserName
	Email *valueobject.Email
}

func NewUser(id int, name *valueobject.UserName, email *valueobject.Email) *User {
	// nilチェック
	return &User{
		Id:    id,
		Name:  name,
		Email: email,
	}
}

func (u *User) Equals(other *User) bool {
	return u.Id == other.Id
}
