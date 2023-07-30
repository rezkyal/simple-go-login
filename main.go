package main

import (
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/rezkyal/simple-go-login/app"
	"github.com/rezkyal/simple-go-login/middlewares"
)

func main() {

	// read config
	cfg, err := app.InitConfig()

	if err != nil {
		log.Fatalf("failed to init config, err: %+v", err)
	}

	// init resources
	resources, err := app.InitResources(cfg)
	if err != nil {
		log.Fatalf("failed to init resources, err: %+v", err)
	}

	// init repos
	repos, err := app.InitRepos(cfg, resources)
	if err != nil {
		log.Fatalf("failed to init repos, err: %+v", err)
	}

	// init repos
	usecase, err := app.InitUsecase(cfg, repos)
	if err != nil {
		log.Fatalf("failed to init uscease, err: %+v", err)
	}

	// init handlers
	handlers, err := app.InitHandlers(usecase)
	if err != nil {
		log.Fatalf("failed to init handlers, err: %+v", err)
	}

	r := gin.Default()

	// return json tag as response validation
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("json")
		})
	}

	public := r.Group("/useraccount")

	public.POST("/signup", handlers.SignupHandler.Signup)
	public.POST("/login", handlers.LoginHandler.Login)

	protected := r.Group("/useraccount/data")
	protected.Use(middlewares.JwtAuthMiddleware(cfg))
	protected.GET("/check", handlers.LoginHandler.IsLoggedIn)

	r.Run(":8080")
}
