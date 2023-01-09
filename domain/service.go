package domain

import "github.com/abassGarane/muscles/domain/models"

type Service interface {
	//workout services
	CreateWorkout(*Workout) (*Workout, error)
	GetWorkout(string) (*Workout, error)
	GetWorkouts(string) ([]*Workout, error)
	DeleteWorkout(string) error
	UpdateWorkout(string, *Workout) (*Workout, error)

	//user services
	GetUserByEmail(string) (*models.User, error)
	CreateUser(*models.UserRequest) (*models.User, error)
	UpdateUser(string, *models.User) (*models.User, error)
}
