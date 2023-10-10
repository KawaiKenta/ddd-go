package valueobject

import (
	"testing"
)

func TestNewUserName(t *testing.T) {
	// 3文字以下でエラーが返ることを確認
	_, err := NewUserName("aa")
	if err == nil {
		t.Error("3文字以下でエラーが返らない")
	}

	// 20文字以上でエラーが返ることを確認
	_, err = NewUserName("aaaaaaaaaaaaaaaaaaaaa")
	if err == nil {
		t.Error("20文字以上でエラーが返らない")
	}

	// 3文字以上20文字以下でエラーが返らないことを確認
	_, err = NewUserName("aaa")
	if err != nil {
		t.Error("3文字以上20文字以下でエラーが返る")
	}
}
