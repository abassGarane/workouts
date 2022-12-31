package domain

type Service interface {
	CreateWorkout(*Workout) (*Workout, error)
	GetWorkout(string) (*Workout, error)
}
