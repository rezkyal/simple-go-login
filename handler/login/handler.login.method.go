package login

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	enUser "github.com/rezkyal/simple-go-login/entity/user"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := context.Background()

	var req enUser.LoginRequest

	err := c.ShouldBindJSON(&req)
	valErrs, ok := err.(validator.ValidationErrors)
	if ok {
		errList := map[string]string{}
		for _, valErr := range valErrs {
			errList[valErr.Field()] = valErr.Tag()
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong input", "fields": errList})
		return
	}

	if err != nil {
		log.Printf("[Login] error when BindJSON %+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResp, err := h.userUsecase.Login(ctx, req)

	if err != nil {
		log.Printf("[Login] error when Login %+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": nil, "success": true, "token": loginResp.Token, "is_password_correct": loginResp.IsPasswordCorrect})
}

func (h *Handler) IsLoggedIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}
