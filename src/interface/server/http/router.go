package http

import (
	"net/http"
	"time"

	"alchemy/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// SetupRouter ...
func SetupRouter(server *echo.Echo, handler *Handler) {

	server.GET("/ping", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "services up and running... "+time.Now().Format(time.RFC3339))
	})
	server.GET("/health", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "OK")
	})

	docs.SwaggerInfo.Schemes = []string{"http"}
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	root := server.Group("/api/v1")
	root.POST("/login", handler.authHandler.login)

	spaceship := root.Group("/spaceship")
	spaceship.POST("", handler.spaceshipHandler.createSpaceship)
	spaceship.PUT("/id/:id", handler.spaceshipHandler.updateSpaceship)
	spaceship.GET("", handler.spaceshipHandler.findSpaceship)
	spaceship.GET("/id/:id", handler.spaceshipHandler.findByID)
}
