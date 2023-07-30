package app

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rezkyal/simple-go-login/entity/config"
	redisPkg "github.com/rezkyal/simple-go-login/pkg/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Resources struct {
	Database *gorm.DB
	Redis    *redisPkg.RedisPkg
}

func InitResources(cfg *config.Config) (*Resources, error) {
	resources := &Resources{}

	// db
	dbCfg := cfg.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", dbCfg.Host, dbCfg.User, dbCfg.Password, dbCfg.Name, dbCfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("[InitResources] failed when init database, err: %+v", err)
	}

	resources.Database = db

	//redis
	redisCli, err := redisPkg.NewRedis(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err != nil {
		return nil, fmt.Errorf("[InitResources] failed when init redis, err: %+v", err)
	}

	resources.Redis = redisCli

	return resources, nil
}
