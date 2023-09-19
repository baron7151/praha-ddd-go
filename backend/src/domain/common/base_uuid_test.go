package domaincommon

import (
	"testing"
)

func TestNewBaseUUID_ValidUUID(t *testing.T) {
	validUUID := "550e8400-e29b-41d4-a716-446655440000"
	id, err := NewBaseUUID(validUUID)
	if err != nil {
		t.Errorf("エラーが発生しました: %v", err)
	}

	if id.GetValue() != validUUID {
		t.Errorf("期待値: %s, 実際の値: %s", validUUID, id.GetValue())
	}
}

func TestNewBaseUUID_InvalidUUID(t *testing.T) {
	invalidUUID := "invalid-uuid"
	_, err := NewBaseUUID(invalidUUID)
	if err == nil {
		t.Error("エラーが発生しませんでしたが、期待されていました")
	}
}

func TestBaseUUID_Equals_SameUUID(t *testing.T) {
	uuid1, _ := NewBaseUUID("")

	if !uuid1.Equals(uuid1) {
		t.Error("UUIDは同じであるはずですが、異なります")
	}
}

func TestBaseUUID_Equals_DifferentUUID(t *testing.T) {
	uuid1, _ := NewBaseUUID("")
	uuid2, _ := NewBaseUUID("")

	if uuid1.Equals(uuid2) {
		t.Error("UUIDは異なるはずですが、同じです")
	}
}
