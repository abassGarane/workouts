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
	workout.CreatedAt = time.Now().Local()
	workout.UpdatedAt = time.Now().Local()
	return s.repo.AddWorkout(workout)
}

func (s *service) GetWorkout(id string) (*Workout, error) {
	return s.repo.GetWorkout(id)
}

func (s *service) GetWorkouts() ([]*Workout, error) {
	return s.repo.GetWorkouts()
}
func (s *service) DeleteWorkout(id string) error {
	return s.repo.DeleteWorkout(id)
}

func (s *service) UpdateWorkout(id string, workout *Workout) (*Workout, error) {
	return s.repo.UpdateWorkout(id, workout)
}
