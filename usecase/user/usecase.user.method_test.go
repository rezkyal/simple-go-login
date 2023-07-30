package user

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/rezkyal/simple-go-login/entity/config"
	enUser "github.com/rezkyal/simple-go-login/entity/user"
)

func TestUsecase_RegisterNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := NewMockUserRepo(ctrl)
	type args struct {
		input enUser.NewUserInput
	}

	dateOfBirthTime, _ := time.Parse("2006-01-02", "2015-02-03")

	tests := []struct {
		name     string
		args     args
		mockFunc func(args)
		wantErr  bool
	}{
		{
			name: "failed to parse time",
			args: args{
				input: enUser.NewUserInput{
					DateOfBirth: "aaaaa",
				},
			},
			mockFunc: func(a args) {},
			wantErr:  true,
		},
		{
			name: "success",
			args: args{
				input: enUser.NewUserInput{
					Email:        "rezky@mail.com",
					FullName:     "Rezky Alamsyah",
					Password:     "passwd",
					PhoneNumber:  "123456789",
					Sex:          "male",
					Biography:    "biography",
					Location:     "location",
					DateOfBirth:  "2015-02-03",
					ProfilePhoto: "img.com",
				},
			},
			mockFunc: func(a args) {
				newUserData := enUser.User{
					Email:        a.input.Email,
					FullName:     a.input.FullName,
					Password:     a.input.Password,
					PhoneNumber:  a.input.PhoneNumber,
					Sex:          a.input.Sex,
					Biography:    a.input.Biography,
					Location:     a.input.Location,
					DateOfBirth:  dateOfBirthTime,
					ProfilePhoto: a.input.ProfilePhoto,
				}

				mockUserRepo.EXPECT().SaveNewUser(gomock.Any(), gomock.Eq(newUserData)).Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := New(&config.Config{}, mockUserRepo)
			tt.mockFunc(tt.args)
			if err := u.RegisterNewUser(context.Background(), tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("Usecase.RegisterNewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecase_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := NewMockUserRepo(ctrl)

	type args struct {
		req enUser.LoginRequest
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func(args)
		want     enUser.LoginResponse
		wantErr  bool
	}{
		{
			name: "error when GetUserData",
			args: args{
				req: enUser.LoginRequest{
					Email:    "rezkyal@mail.com",
					Password: "aaa",
				},
			},
			mockFunc: func(a args) {
				mockUserRepo.EXPECT().GetUserData(gomock.Any(), gomock.Eq(a.req.Email)).Return(enUser.User{}, errors.New("error test"))
			},
			want:    enUser.LoginResponse{},
			wantErr: true,
		},
		{
			name: "error when CompareHashAndPassword",
			args: args{
				req: enUser.LoginRequest{
					Email:    "rezkyal@mail.com",
					Password: "aaa",
				},
			},
			mockFunc: func(a args) {
				mockUserRepo.EXPECT().GetUserData(gomock.Any(), gomock.Eq(a.req.Email)).Return(enUser.User{
					Password: "aaa",
				}, nil)
			},
			want:    enUser.LoginResponse{},
			wantErr: true,
		},
		{
			name: "password mismatch",
			args: args{
				req: enUser.LoginRequest{
					Email:    "rezkyal@mail.com",
					Password: "password",
				},
			},
			mockFunc: func(a args) {
				mockUserRepo.EXPECT().GetUserData(gomock.Any(), gomock.Eq(a.req.Email)).Return(enUser.User{
					Password: "$2a$10$1JE/Iux083fzPDoSVa4Wx.tUsThgaB3ZuZVDMfUu6XXAZQAmrK6q6",
				}, nil)
			},
			want:    enUser.LoginResponse{},
			wantErr: false,
		},
		{
			name: "success",
			args: args{
				req: enUser.LoginRequest{
					Email:    "rezkyal@mail.com",
					Password: "password",
				},
			},
			mockFunc: func(a args) {
				mockUserRepo.EXPECT().GetUserData(gomock.Any(), gomock.Eq(a.req.Email)).Return(enUser.User{
					ID:       10,
					Password: "$2a$10$E1olSV1JXl/i2A46z0NFW.6m0TTgih6fWBI85Xu2EubUU3xZGexFm",
				}, nil)
			},
			want: enUser.LoginResponse{
				Token:             "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTA3NDIzMDAsInVzZXJfaWQiOjEwfQ.lETn2HXS3I0Q__IZVNr3pJ65dJrJRh0ECJCzCLAD1iU",
				IsPasswordCorrect: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(tt.args)
			u, _ := New(&config.Config{
				Token: config.Token{
					Lifespan: 100,
					Secret:   "secret",
				},
			}, mockUserRepo)
			got, err := u.Login(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.IsPasswordCorrect, tt.want.IsPasswordCorrect) {
				t.Errorf("Usecase.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
