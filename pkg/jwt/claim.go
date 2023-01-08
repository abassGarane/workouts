package jwt

import "github.com/golang-jwt/jwt/v4"

type Claim struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Admin    bool   `json:"admin"`
	jwt.RegisteredClaims
}
