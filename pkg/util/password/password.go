package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	convert_type "nam_0508/pkg/util/convert-type"
)

const cost = 12

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(convert_type.StringToBytes(password), cost)
	return convert_type.BytesToString(bytes), err
}

func CheckPasswordHash(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword(convert_type.StringToBytes(hashedPassword), convert_type.StringToBytes(password))
	return err == nil
}

func ValidatePassword(p string) error {
	if len(p) > 100 || len(p) < 8 {
		return errors.New("invalid password")
	}
	return nil
}
