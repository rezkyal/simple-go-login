package user

import (
	"github.com/rezkyal/simple-go-login/entity/config"
	"gorm.io/gorm"
)

type Repo struct {
	cfg   *config.Config
	db    *gorm.DB
	redis Redis
}

func New(cfg *config.Config, db *gorm.DB, redis Redis) (*Repo, error) {
	return &Repo{
		cfg:   cfg,
		db:    db,
		redis: redis,
	}, nil
}
