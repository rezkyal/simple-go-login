package login

import (
	"context"

	enUser "github.com/rezkyal/simple-go-login/entity/user"
)

//go:generate mockgen -source=handler.login.deps.go -destination=handler.login.deps_mock.go -package=login
type UserUsecase interface {
	Login(ctx context.Context, req enUser.LoginRequest) (enUser.LoginResponse, error)
}
