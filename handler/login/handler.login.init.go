package login

type Handler struct {
	userUsecase UserUsecase
}

func New(userUsecase UserUsecase) (*Handler, error) {
	return &Handler{userUsecase: userUsecase}, nil
}
