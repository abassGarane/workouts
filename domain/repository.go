package domain

type Repository interface {
	AddWorkout(*Workout) (*Workout, error)
	GetWorkout(string) (*Workout, error)
}
