package user

import (
	"context"
	"fmt"
	"time"

	enUser "github.com/rezkyal/simple-go-login/entity/user"
	"github.com/rezkyal/simple-go-login/utils"
)

func (u *Usecase) RegisterNewUser(ctx context.Context, input enUser.NewUserInput) error {
	dateOfBirth, err := time.Parse("2006-01-02", input.DateOfBirth)

	if err != nil {
		return fmt.Errorf("[SaveNewUser] error when time.Parse, err: %+v", err)
	}

	newUserData := enUser.User{
		Email:        input.Email,
		FullName:     input.FullName,
		Password:     input.Password,
		PhoneNumber:  input.PhoneNumber,
		Sex:          input.Sex,
		Biography:    input.Biography,
		Location:     input.Location,
		DateOfBirth:  dateOfBirth,
		ProfilePhoto: input.ProfilePhoto,
	}

	return u.userRepo.SaveNewUser(ctx, newUserData)
}

func (u *Usecase) Login(ctx context.Context, req enUser.LoginRequest) (enUser.LoginResponse, error) {
	var resp enUser.LoginResponse = enUser.LoginResponse{}
	userData, err := u.userRepo.GetUserData(ctx, req.Email)

	if err != nil {
		return resp, fmt.Errorf("[Login] error when GetUserData, err: %+v", err)
	}

	isPasswordCorrect, err := utils.VerifyPassword(userData.Password, req.Password)

	if err != nil {
		return resp, fmt.Errorf("[Login] error when VerifyPassword, err: %+v", err)
	}

	if !isPasswordCorrect {
		return resp, nil
	}

	resp.IsPasswordCorrect = true

	token, err := utils.GenerateToken(userData.ID, u.cfg.Token.Lifespan, u.cfg.Token.Secret)

	if err != nil {
		return resp, fmt.Errorf("[Login] error when GenerateToken, err: %+v", err)
	}

	resp.Token = token

	return enUser.LoginResponse{
		IsPasswordCorrect: true,
		Token:             token,
	}, nil
}
