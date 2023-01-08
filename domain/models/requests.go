package models

import "github.com/go-playground/validator/v10"

type UserRequest struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u *UserRequest) Validate() error {
	val := validator.New()
	return val.Struct(u)
}
