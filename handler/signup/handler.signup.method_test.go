package signup

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"

	enUser "github.com/rezkyal/simple-go-login/entity/user"
)

func TestHandler_Signup(t *testing.T) {
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
			responseBody: `{"error":"wrong input","fields":{"Biography":"required","DateOfBirth":"required","Email":"required","FullName":"required","Location":"required","Password":"required","PhoneNumber":"required","ProfilePhoto":"required","Sex":"required"}}`,
			statusCode:   http.StatusBadRequest,
		},
		{
			name: "error on RegisterNewUser",
			mockFunc: func() {
				req := enUser.NewUserInput{
					Biography:    "testing",
					DateOfBirth:  "1990-03-28",
					Email:        "rezkyal@mail.com",
					FullName:     "Rezky Alamsyah",
					Location:     "Balikpapan",
					Password:     "password",
					PhoneNumber:  "123456789",
					ProfilePhoto: "http://img.profile",
					Sex:          "male",
				}
				mockUserUsecase.EXPECT().RegisterNewUser(gomock.Any(), gomock.Eq(req)).Return(errors.New("error test"))
			},
			requestBody: `{
				"biography": "testing",
				"date_of_birth": "1990-03-28",
				"email": "rezkyal@mail.com",
				"fullname": "Rezky Alamsyah",
				"location": "Balikpapan",
				"password": "password",
				"phone_number": "123456789",
				"profile_photo": "http://img.profile",
				"sex": "male"    
			}`,
			responseBody: `{"error":"error test"}`,
			statusCode:   http.StatusInternalServerError,
		},
		{
			name: "success",
			mockFunc: func() {
				req := enUser.NewUserInput{
					Biography:    "testing",
					DateOfBirth:  "1990-03-28",
					Email:        "rezkyal@mail.com",
					FullName:     "Rezky Alamsyah",
					Location:     "Balikpapan",
					Password:     "password",
					PhoneNumber:  "123456789",
					ProfilePhoto: "http://img.profile",
					Sex:          "male",
				}
				mockUserUsecase.EXPECT().RegisterNewUser(gomock.Any(), gomock.Eq(req)).Return(nil)
			},
			requestBody: `{
				"biography": "testing",
				"date_of_birth": "1990-03-28",
				"email": "rezkyal@mail.com",
				"fullname": "Rezky Alamsyah",
				"location": "Balikpapan",
				"password": "password",
				"phone_number": "123456789",
				"profile_photo": "http://img.profile",
				"sex": "male"    
			}`,
			responseBody: `{"error":null,"success":true}`,
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

			h.Signup(c)

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
