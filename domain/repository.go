package domain

import "github.com/abassGarane/muscles/domain/models"

type Repository interface {
	// Workouts
	AddWorkout(*Workout) (*Workout, error)
	GetWorkout(string) (*Workout, error)
	GetWorkouts(string) ([]*Workout, error)
	DeleteWorkout(string) error
	UpdateWorkout(string, *Workout) (*Workout, error)

	//Users
	CreateUser(*models.User) (*models.User, error)
	UpdateUser(string, *models.User) (*models.User, error) // by email
	GetUserByEmail(string) (*models.User, error)
	// DeleteUser(string) error
}
