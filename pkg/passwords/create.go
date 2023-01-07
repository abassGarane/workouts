package passwords

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreateHashedPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not hash password %v", err))
	}
	fmt.Println(pass)
	return string(pass), nil
}
