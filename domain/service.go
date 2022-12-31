package domain

type Service interface {
	CreateWorkout(*Workout) (*Workout, error)
}
