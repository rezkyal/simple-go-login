package app

import (
	"fmt"

	"github.com/rezkyal/simple-go-login/entity/config"
	userUsecase "github.com/rezkyal/simple-go-login/usecase/user"
)

type Usecases struct {
	UserUsecase *userUsecase.Usecase
}

func InitUsecase(cfg *config.Config, repos *Repos) (*Usecases, error) {
	usecase := &Usecases{}

	userU, err := userUsecase.New(cfg, repos.UserRepo)
	if err != nil {
		return nil, fmt.Errorf("[InitUsecase] init userusecase error, err: %+v", err)
	}

	usecase.UserUsecase = userU

	return usecase, nil
}
