package valueobject

import (
	"errors"
)

type UserName struct {
	value string
}

func NewUserName(value string) (*UserName, error) {
	// validation 3文字以上20文字以下
	if len(value) < 3 || len(value) > 20 {
		return nil, errors.New("ユーザ名は3文字以上20文字以下で入力してください")
	}
	return &UserName{
		value: value,
	}, nil
}

func (u *UserName) Value() string {
	return u.value
}

func (u *UserName) Equals(other *UserName) bool {
	return u.value == other.value
}
