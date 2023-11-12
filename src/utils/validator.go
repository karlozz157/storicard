package utils

import (
	"regexp"

	"github.com/karlozz157/storicard/src/domain/errors"
)

const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func ValidateEmail(email string) error {
	regex := regexp.MustCompile(emailPattern)
	isValid := regex.MatchString(email)

	if isValid {
		return nil
	}

	return errors.ErrEmailInvalid
}
