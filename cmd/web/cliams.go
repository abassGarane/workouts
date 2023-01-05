package main

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}
