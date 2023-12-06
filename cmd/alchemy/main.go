package main

import (
	"alchemy/src/interface/container"
	"alchemy/src/interface/server"
)

// @title    alchemy
// @version  1.0
// @schemes http https
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	server.StartService(container.Setup())
}
