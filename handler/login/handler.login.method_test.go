package login

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	enUser "github.com/rezkyal/simple-go-login/entity/user"
)

func TestHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserUsecase := NewMockUserUsecase(ctrl)
	tests := []struct {
		name         string
		mockFunc     func()
		requestBody  string
		responseBody string
		statusCode   int
	}{
		{
			name:         "error empty request",
			mockFunc:     func() {},
			requestBody:  "",
			responseBody: `{"error":"EOF"}`,
			statusCode:   http.StatusBadRequest,
		},
		{
			name:         "error when bind json",
			mockFunc:     func() {},
			requestBody:  "aaaaa",
			responseBody: `{"error":"invalid character 'a' looking for beginning of value"}`,
			statusCode:   http.StatusBadRequest,
		},
		{
			name:         "error validation",
			mockFunc:     func() {},
			requestBody:  "{}",
			responseBody: `{"error":"wrong input","fields":{"Email":"required","Password":"required"}}`,
			statusCode:   http.StatusBadRequest,
		},
		{
			name: "error on Login",
			mockFunc: func() {
				req := enUser.LoginRequest{
					Email:    "rezkyal@mail.com",
					Password: "password",
				}
				mockUserUsecase.EXPECT().Login(gomock.Any(), gomock.Eq(req)).Return(enUser.LoginResponse{}, errors.New("error test"))
			},
			requestBody: `{
				"email": "rezkyal@mail.com",
				"password": "password"
			}`,
			responseBody: `{"error":"error test"}`,
			statusCode:   http.StatusInternalServerError,
		},
		{
			name: "success",
			mockFunc: func() {
				req := enUser.LoginRequest{
					Email:    "rezkyal@mail.com",
					Password: "password",
				}
				mockUserUsecase.EXPECT().Login(gomock.Any(), gomock.Eq(req)).Return(enUser.LoginResponse{
					Token:             "tokentoken",
					IsPasswordCorrect: true,
				}, nil)
			},
			requestBody: `{
				"email": "rezkyal@mail.com",
				"password": "password"
			}`,
			responseBody: `{"error":null,"is_password_correct":true,"success":true,"token":"tokentoken"}`,
			statusCode:   http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, _ := New(mockUserUsecase)
			tt.mockFunc()

			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)
			c.Request = &http.Request{
				Header: make(http.Header),
			}

			c.Request.Method = "POST"
			c.Request.Header.Set("Content-type", "application/json")

			c.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(tt.requestBody)))

			h.Login(c)

			if w.Code != tt.statusCode {
				t.Errorf("wrong statusCode, expected = %+v, actual = %+v", tt.statusCode, w.Code)
				return
			}

			respBodyRaw, _ := ioutil.ReadAll(w.Body)
			respBody := string(respBodyRaw)

			if respBody != tt.responseBody {
				t.Errorf("wrong responseBody, expected = %+v, actual = %+v", tt.responseBody, respBody)
				return
			}

		})
	}
}

func TestHandler_IsLoggedIn(t *testing.T) {
	tests := []struct {
		name         string
		requestBody  string
		responseBody string
		statusCode   int
	}{
		{
			name:         "success",
			requestBody:  "",
			responseBody: `{"success":true}`,
			statusCode:   http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{}

			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)
			c.Request = &http.Request{
				Header: make(http.Header),
			}

			c.Request.Method = "POST"
			c.Request.Header.Set("Content-type", "application/json")

			c.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(tt.requestBody)))

			h.IsLoggedIn(c)

			if w.Code != tt.statusCode {
				t.Errorf("wrong statusCode, expected = %+v, actual = %+v", tt.statusCode, w.Code)
				return
			}

			respBodyRaw, _ := ioutil.ReadAll(w.Body)
			respBody := string(respBodyRaw)

			if respBody != tt.responseBody {
				t.Errorf("wrong responseBody, expected = %+v, actual = %+v", tt.responseBody, respBody)
				return
			}
		})
	}
}
