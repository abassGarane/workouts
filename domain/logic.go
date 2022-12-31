package domain

import (
	"time"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateWorkout(workout *Workout) (*Workout, error) {
	workout.CreatedAt = time.Now()
	workout.UpdatedAt = time.Now()
	return s.repo.AddWorkout(workout)
}
