package main

import (
	"nguyendangviet12022004/url-short/url/internal/handler"
	"nguyendangviet12022004/url-short/url/internal/model"
	"nguyendangviet12022004/url-short/url/internal/repository"
	"nguyendangviet12022004/url-short/url/internal/route"
	"nguyendangviet12022004/url-short/url/internal/service"
	"nguyendangviet12022004/url-short/url/pkg/config"
	"nguyendangviet12022004/url-short/url/pkg/database"
	"nguyendangviet12022004/url-short/url/pkg/logging"
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
	db.AutoMigrate(&model.Url{})

	engine := gin.Default()

	// dependency injection
	urlRepository := repository.NewUrlRepository(db)
	urlService := service.NewUrlService(urlRepository)
	urlHandler := handler.NewUrlHandler(urlService)

	// route
	route.InitUrlRoute(engine, urlHandler)

	engine.Run(":" + config.Server.Port)
}
