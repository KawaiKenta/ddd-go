package valueobject

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	// validation regex email形式
	regexpEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,6}$`)
	if !regexpEmail.MatchString(value) {
		return nil, errors.New("メールアドレスの形式が不正です")
	}
	return &Email{
		value: value,
	}, nil
}

func (e *Email) Value() string {
	return e.value
}
