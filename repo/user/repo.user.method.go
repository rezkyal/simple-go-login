package user

import (
	"context"
	"fmt"
	"time"

	enUser "github.com/rezkyal/simple-go-login/entity/user"

	"golang.org/x/crypto/bcrypt"
)

func (r *Repo) SaveNewUser(ctx context.Context, input enUser.User) error {

	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), r.cfg.Password.Cost)
	if err != nil {
		return fmt.Errorf("[SaveNewUser] error when hash password, err: %+v", err)
	}

	input.Password = string(hashedPassword)

	result := r.db.Create(&input)

	if result.Error != nil {
		return fmt.Errorf("[SaveNewUser] error when create new user, err: %+v", result.Error)
	}

	return nil
}

func (r *Repo) GetUserData(ctx context.Context, email string) (enUser.User, error) {
	var err error

	u := enUser.User{}

	err = r.db.Model(enUser.User{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		return u, fmt.Errorf("[GetUserData] error when query the user data, err: %+v", err)
	}

	return u, nil
}
