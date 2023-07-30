package user

import (
	"github.com/rezkyal/simple-go-login/entity/config"
	"gorm.io/gorm"
)

type Repo struct {
	cfg *config.Config
	db  *gorm.DB
}

func New(cfg *config.Config, db *gorm.DB) (*Repo, error) {
	return &Repo{
		cfg: cfg,
		db:  db,
	}, nil
}
