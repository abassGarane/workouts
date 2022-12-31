package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workout struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Type      string             `json:"type" bson:"type,omitempty" validate:"required"`
	Reps      int                `json:"reps,string" bson:"reps,omitempty" validate:"required"`
	Load      int                `json:"load,string" bson:"load,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty" validate:"required, datetime"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty" validate:"required, datetime"`
}

//TODO: implement validate for presence of all fields in struct
