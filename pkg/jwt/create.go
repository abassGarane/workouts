package jwt

import "github.com/golang-jwt/jwt/v4"

func CreateSignature(claims *Claim, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString([]byte(secret))
	return tokenStr
}
