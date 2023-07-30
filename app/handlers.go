package app

import (
	"fmt"

	loginHandler "github.com/rezkyal/simple-go-login/handler/login"
	signupHandler "github.com/rezkyal/simple-go-login/handler/signup"
)

type Handlers struct {
	SignupHandler *signupHandler.Handler
	LoginHandler  *loginHandler.Handler
}

func InitHandlers(usecases *Usecases) (*Handlers, error) {
	handlers := &Handlers{}

	// signup handlers
	signupH, err := signupHandler.New(usecases.UserUsecase)

	if err != nil {
		return handlers, fmt.Errorf("[InitHandler] error init signup handler, err: %+v", err)
	}

	handlers.SignupHandler = signupH

	//loign handlers
	loginH, err := loginHandler.New(usecases.UserUsecase)

	if err != nil {
		return handlers, fmt.Errorf("InitHandler error init login handler, err: %+v", err)
	}

	handlers.LoginHandler = loginH

	return handlers, nil
}
