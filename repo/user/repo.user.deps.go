package user

import (
	"context"
	"time"
)

//go:generate mockgen -source=repo.user.deps.go -destination=repo.user.deps_mock.go -package=user
type Redis interface {
	SetJSONTTL(ctx context.Context, key string, value interface{}, TTL time.Duration) error
	GetJSON(ctx context.Context, key string, data interface{}) error
}
