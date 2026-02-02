package main

import (
	"net/http"
	"nguyendangviet12022004/url-short/notification/pkg/logging"
	"nguyendangviet12022004/url-short/notification/pkg/websocket"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	// init logger
	err := logging.InitLogger("development")
	if err != nil {
		logging.GetLogger().Fatal("Failed to initialize logger: ", zap.Error(err))
	}
	defer logging.GetLogger().Sync()

	// init hub
	hub := websocket.NewHub()
	go hub.Run()

	// init gin
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	// init route
	g.GET("/ws", func(c *gin.Context) {
		userId := c.GetHeader("X-User-Id")

		if userId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}
		websocket.ServeWs(uint(userIdInt), hub, c.Writer, c.Request)
	})

	// run server
	g.Run(":8080")
}
