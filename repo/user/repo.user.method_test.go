package user

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/rezkyal/simple-go-login/entity/config"
	enRedis "github.com/rezkyal/simple-go-login/entity/redis"
	enUser "github.com/rezkyal/simple-go-login/entity/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRepo_SaveNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRedis := NewMockRedis(ctrl)

	dbMock, mockDB, _ := sqlmock.New()

	defer dbMock.Close()

	db, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: dbMock,
	}), &gorm.Config{})

	type fields struct {
		cfg *config.Config
	}
	type args struct {
		input enUser.User
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockFunc func(args)
		wantErr  bool
	}{
		{
			name: "failed GenerateFromPassword",
			fields: fields{
				cfg: &config.Config{
					Password: config.Password{
						Cost: 10,
					},
				},
			},
			args: args{
				input: enUser.User{
					Password: "passssssssssasfsafsacsalckasnmclsakcmnasolckmasocklasmcoaslkcnmsaolckasmnoclkiasmcas",
				},
			},
			mockFunc: func(a args) {},
			wantErr:  true,
		},
		{
			name: "failed When query",
			fields: fields{
				cfg: &config.Config{
					Password: config.Password{
						Cost: 10,
					},
				},
			},
			args: args{
				input: enUser.User{
					Email:    "testemail",
					Password: "pass",
				},
			},
			mockFunc: func(a args) {
				mockDB.ExpectBegin()
				mockDB.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "useraccount" ("email","fullname","password","phone_number","sex","biography","location","date_of_birth","profile_photo","refresh_token","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING "id"`)).
					WithArgs("testemail", sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("error test"))
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				cfg: &config.Config{
					Password: config.Password{
						Cost: 10,
					},
				},
			},
			args: args{
				input: enUser.User{
					Email:    "testemail1",
					FullName: "fullname",
					Password: "pass",
				},
			},
			mockFunc: func(a args) {
				mockDB.ExpectBegin()
				mockDB.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "useraccount" ("email","fullname","password","phone_number","sex","biography","location","date_of_birth","profile_photo","refresh_token","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING "id"`)).
					WithArgs("testemail1", sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mockDB.ExpectCommit()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := New(tt.fields.cfg, db, mockRedis)
			tt.mockFunc(tt.args)
			if err := r.SaveNewUser(context.Background(), tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("Repo.SaveNewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_GetUserData(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRedis := NewMockRedis(ctrl)

	dbMock, mockDB, _ := sqlmock.New()

	defer dbMock.Close()

	db, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: dbMock,
	}), &gorm.Config{})

	type fields struct {
		cfg *config.Config
	}
	type args struct {
		email string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockFunc func(args)
		want     enUser.User
		wantErr  bool
	}{
		{
			name: "success from redis",
			fields: fields{
				cfg: &config.Config{},
			},
			args: args{
				email: "mail",
			},
			mockFunc: func(a args) {
				var (
					redisKey = fmt.Sprintf(enRedis.KeyGetUserData, a.email)
				)

				u := enUser.User{
					ID:       12,
					Password: "password",
				}

				mockRedis.EXPECT().GetJSON(gomock.Any(), gomock.Eq(redisKey), gomock.Any()).Return(nil).SetArg(2, u)
			},
			want: enUser.User{
				ID:       12,
				Password: "password",
			},
			wantErr: false,
		},
		{
			name: "error when query",
			fields: fields{
				cfg: &config.Config{},
			},
			args: args{
				email: "mail",
			},
			mockFunc: func(a args) {
				var (
					redisKey = fmt.Sprintf(enRedis.KeyGetUserData, a.email)
				)

				mockRedis.EXPECT().GetJSON(gomock.Any(), gomock.Eq(redisKey), gomock.Any()).Return(errors.New("error test"))
				mockDB.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "useraccount" WHERE email = $1 LIMIT 1`)).
					WithArgs("mail").
					WillReturnError(errors.New("error test"))
			},
			want:    enUser.User{},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				cfg: &config.Config{},
			},
			args: args{
				email: "mail",
			},
			mockFunc: func(a args) {
				var (
					redisKey = fmt.Sprintf(enRedis.KeyGetUserData, a.email)
				)

				mockRedis.EXPECT().GetJSON(gomock.Any(), gomock.Eq(redisKey), gomock.Any()).Return(errors.New("error test"))
				mockDB.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "useraccount" WHERE email = $1 LIMIT 1`)).
					WithArgs("mail").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mockRedis.EXPECT().SetJSONTTL(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			want: enUser.User{
				ID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := New(tt.fields.cfg, db, mockRedis)
			tt.mockFunc(tt.args)
			got, err := r.GetUserData(context.Background(), tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.GetUserData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repo.GetUserData() = %v, want %v", got, tt.want)
			}
		})
	}
}
