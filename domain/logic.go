package domain

import (
	"log"
	"time"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateWorkout(workout *Workout) (*Workout, error) {
	nai, err := time.LoadLocation("Africa/Nairobi")
	if err != nil {
		log.Fatal(err)
	}
	workout.CreatedAt = time.Now().UTC().Local().In(nai)
	workout.UpdatedAt = time.Now().UTC().Local().In(nai)
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
	workout.UpdatedAt = time.Now().UTC().Local()
	return s.repo.UpdateWorkout(id, workout)
}
