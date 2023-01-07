package models

import (
	"github.com/abassGarane/muscles/pkg/passwords"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"id,string" bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Email          string             `json:"email" bson:"email" validate:"required,email"`
	HashedPassword string             `json:"hashed_password" bson:"hashed_password,omitempty" validate:"required"`
	HomeAddress    *Address           `json:"home_address" bson:"home_address" validate:"required"`
}

func NewUser(name, email, password, county, town, village string) *User {
	pass, _ := passwords.CreateHashedPassword(password)
	return &User{
		Name:           name,
		Email:          email,
		HashedPassword: pass,
		HomeAddress: &Address{
			HomeCounty:  county,
			HomeTown:    town,
			HomeVillage: village,
		},
	}
}
func (s *User) Validate() error {
	val := validator.New()
	return val.Struct(s)
}
