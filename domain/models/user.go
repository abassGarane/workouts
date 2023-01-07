package models

import (
	"time"

	"github.com/abassGarane/muscles/pkg/passwords"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"id,string" bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Email          string             `json:"email" bson:"email" validate:"required,email"`
	HashedPassword string             `json:"hashed_password" bson:"hashed_password,omitempty" validate:"required"`
	Admin          bool               `json:"admin" bson:"admin" validate:"-"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at,omitempty" `
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

func NewUser(name, email, password string) *User {
	pass, _ := passwords.CreateHashedPassword(password)
	return &User{
		Name:           name,
		Email:          email,
		HashedPassword: pass,
	}
}
func (s *User) Validate() error {
	val := validator.New()
	return val.Struct(s)
}
