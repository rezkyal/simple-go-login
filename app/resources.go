package app

import (
	"fmt"

	"github.com/rezkyal/simple-go-login/entity/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Resources struct {
	Database *gorm.DB
}

func InitResources(cfg *config.Config) (*Resources, error) {
	resources := &Resources{}

	dbCfg := cfg.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", dbCfg.Host, dbCfg.User, dbCfg.Password, dbCfg.Name, dbCfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("[InitResources] failed when init database, err: %+v", err)
	}

	resources.Database = db

	return resources, nil
}
