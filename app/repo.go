package app

import (
	"fmt"

	"github.com/rezkyal/simple-go-login/entity/config"
	userRepo "github.com/rezkyal/simple-go-login/repo/user"
)

type Repos struct {
	UserRepo *userRepo.Repo
}

func InitRepos(cfg *config.Config, resources *Resources) (*Repos, error) {
	repos := &Repos{}

	userR, err := userRepo.New(cfg, resources.Database, resources.Redis)
	if err != nil {
		return nil, fmt.Errorf("[InitRepo] failed to init user repo %+v", err)
	}

	repos.UserRepo = userR

	return repos, nil
}
