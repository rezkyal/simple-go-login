package login

import (
	"context"

	enUser "github.com/rezkyal/simple-go-login/entity/user"
)

type UserUsecase interface {
	Login(ctx context.Context, req enUser.LoginRequest) (enUser.LoginResponse, error)
}
