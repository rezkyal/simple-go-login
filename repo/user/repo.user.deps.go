package user

import (
	"context"
	"time"
)

type Redis interface {
	SetJSONTTL(ctx context.Context, key string, value interface{}, TTL time.Duration) error
	GetJSON(ctx context.Context, key string, data interface{}) error
}
