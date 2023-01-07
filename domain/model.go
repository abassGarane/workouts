package domain

import (
	"time"

	validator "github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workout struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" `
	Type      string             `json:"type" bson:"type,omitempty" validate:"required"`
	Reps      int                `json:"reps,string" bson:"reps,omitempty" validate:"required"`
	Load      int                `json:"load,string" bson:"load,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty" `
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

// TODO: implement validate for presence of all fields in struct
// TODO: Add generics to handle different structs
func Validate[S any](w S) error {
	val := validator.New()
	return val.Struct(w)
}
