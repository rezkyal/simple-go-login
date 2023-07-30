package signup

import (
	"context"

	enUser "github.com/rezkyal/simple-go-login/entity/user"
)

type UserUsecase interface {
	RegisterNewUser(ctx context.Context, input enUser.NewUserInput) error
}
