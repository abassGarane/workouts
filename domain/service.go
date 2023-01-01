package domain

type Service interface {
	CreateWorkout(*Workout) (*Workout, error)
	GetWorkout(string) (*Workout, error)
	GetWorkouts() ([]*Workout, error)
	DeleteWorkout(string)error
	UpdateWorkout(string, *Workout)(*Workout, error)
}
