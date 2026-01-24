package main

import (
	"nguyendangviet12022004/url-short/user/internal/handler"
	"nguyendangviet12022004/url-short/user/internal/model"
	"nguyendangviet12022004/url-short/user/internal/repository"
	"nguyendangviet12022004/url-short/user/internal/route"
	"nguyendangviet12022004/url-short/user/internal/service"
	"nguyendangviet12022004/url-short/user/pkg/config"
	"nguyendangviet12022004/url-short/user/pkg/database"
	"nguyendangviet12022004/url-short/user/pkg/jwt"
	"nguyendangviet12022004/url-short/user/pkg/logging"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	// load env
	var env = os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	// init logger
	err := logging.InitLogger(env)
	if err != nil {
		panic(err)
	}

	// load config
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// init validator
	validator.New(validator.WithRequiredStructEnabled())

	// connect db
	db, err := database.ConnectDb(&config.Db)
	if err != nil {
		panic(err)
	}

	// migrate db & index
	db.AutoMigrate(&model.User{})
	db.Migrator().CreateIndex(&model.User{}, "email")

	// load private key
	err = jwt.LoadPrivateKey(config.Jwt.PrivateKeyPath)
	if err != nil {
		panic(err)
	}

	// di
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// init engine
	engine := gin.Default()

	// init route
	route.InitUserRoute(engine, userHandler)

	engine.Run(":" + config.Server.Port)

}
