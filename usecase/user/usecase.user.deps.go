package user

import (
	"context"

	enUser "github.com/rezkyal/simple-go-login/entity/user"
)

//go:generate mockgen -source=usecase.user.deps.go -destination=usecase.user.deps_mock.go -package=user
type UserRepo interface {
	SaveNewUser(ctx context.Context, input enUser.User) error
	GetUserData(ctx context.Context, email string) (enUser.User, error)
}
