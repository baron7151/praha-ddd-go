package domaincommon

import (
	"regexp"

	"github.com/google/uuid"
)

var UUIDPattern = regexp.MustCompile(`^([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})$`)

type BaseUUID struct {
	value string
}

// type UUIDProvider interface {
// 	GetValue() string
// 	Equals(other UUIDProvider) bool
// }

func NewBaseUUID(value string) (BaseUUID, error) {
	if value == "" {
		value = uuid.New().String()
	} else if !UUIDPattern.MatchString(value) {
		return BaseUUID{}, NewDomainError("This ID is invalid ${value}")

	}
	return BaseUUID{value: value}, nil
}

func (u BaseUUID) GetValue() string {
	return u.value
}

func (u BaseUUID) Equals(b BaseUUID) bool {
	return u.value == b.value

}
