package user

import "github.com/rezkyal/simple-go-login/entity/config"

type Usecase struct {
	cfg      *config.Config
	userRepo UserRepo
}

func New(cfg *config.Config, userRepo UserRepo) (*Usecase, error) {
	return &Usecase{
		cfg:      cfg,
		userRepo: userRepo,
	}, nil
}
