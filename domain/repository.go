package domain

type Repository interface {
	AddWorkout(*Workout) (*Workout, error)
	GetWorkout(string) (*Workout, error)
	GetWorkouts() ([]*Workout, error)
	DeleteWorkout(string)error
	UpdateWorkout(string, *Workout)(*Workout, error)
}
