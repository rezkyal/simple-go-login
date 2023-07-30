package signup

import (
	"context"

	enUser "github.com/rezkyal/simple-go-login/entity/user"
)

//go:generate mockgen -source=handler.signup.deps.go -destination=handler.signup.deps_mock.go -package=signup
type UserUsecase interface {
	RegisterNewUser(ctx context.Context, input enUser.NewUserInput) error
}
