package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	public := r.Group("/useraccount")

	public.POST("/signup", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})
	})

	r.Run(":8080")
}
