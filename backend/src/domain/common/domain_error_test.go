package domaincommon

import (
	"reflect"
	"testing"
)

func TestNewDomainError(t *testing.T) {
	message := "Domain Errorが発生しました。"
	err := NewDomainError(message)
	if reflect.TypeOf(err).String() != "domaincommon.DomainError" {
		t.Errorf("期待値: *domaincommon.DomainError, 実際の値: %s", reflect.TypeOf(err).String())
	}
}

func TestError(t *testing.T) {
	message := "Domain Errorが発生しました。"
	err := NewDomainError(message)
	if err.Error() != message {
		t.Errorf("期待値: %s, 実際の値: %s", message, err.Error())
	}
}
