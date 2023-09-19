package domaincommon

import (
	"testing"
)

func TestNewEmail_ValidEmail(t *testing.T) {
	validEmail := "test1@example.com"
	email, err := NewEmail(validEmail)
	if err != nil {
		t.Errorf("エラーが発生しました: %v", err)
	}
	if email.value != validEmail {
		t.Errorf("期待値: %s, 実際の値: %s", validEmail, email.value)
	}
}

func TestNewEmail_InValidEmail(t *testing.T) {
	invalidEmail := "invalid-email"
	_, err := NewEmail(invalidEmail)
	if err == nil {
		t.Error("エラーが発生しませんでしたが、エラーを期待していました。")
	}
}

func TestEquals_Equal(t *testing.T) {
	email1, _ := NewEmail("test1@example.com")
	email2, _ := NewEmail("test1@example.com")
	result := email1.Equals(email2)
	if !result {
		t.Error("メールアドレスは同じであるはずですが、異なります")
	}
}
