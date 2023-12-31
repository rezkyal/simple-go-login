package signup

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	enUser "github.com/rezkyal/simple-go-login/entity/user"
)

func (h *Handler) Signup(c *gin.Context) {
	ctx := context.Background()
	var input enUser.NewUserInput

	err := c.ShouldBindJSON(&input)

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
		log.Printf("[Signup] error when BindJSON %+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.userUsecase.RegisterNewUser(ctx, input)

	if err != nil {
		log.Printf("[Signup] error when RegisterNewUser %+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": nil, "success": true})
}
