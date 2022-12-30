package domain

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}
