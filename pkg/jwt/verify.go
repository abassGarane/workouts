package jwt

import "golang.org/x/crypto/bcrypt"

func Verify(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
