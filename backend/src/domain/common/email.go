package domaincommon

import "regexp"

var EmailPattern = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-.]+\.[a-z]+\z`)

type Email struct {
	value string
}

func ValidateEmail(value string) bool {
	if value == "" {
		return false
	} else if !EmailPattern.MatchString(value) {
		return false
	} else {
		return true
	}
}

func NewEmail(value string) (Email, error) {
	if !ValidateEmail(value) {
		return Email{}, NewDomainError("This email is invalid ${value}")
	}
	return Email{value: value}, nil
}

func (e Email) GetValue() string {
	return e.value
}

func (e Email) Equals(other Email) bool {
	return e.value == other.value
}
