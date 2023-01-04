package domain

import (
	"time"
)

type Workout struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Type      string    `json:"type" bson:"type,omitempty" validate:"required"`
	Reps      int       `json:"reps,string" bson:"reps,omitempty" validate:"required"`
	Load      int       `json:"load,string" bson:"load,omitempty" validate:"required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty" validate:"required, datetime"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty" validate:"required, datetime"`
}

//TODO: implement validate for presence of all fields in struct
