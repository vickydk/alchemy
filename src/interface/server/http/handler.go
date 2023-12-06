package http

import "alchemy/src/interface/container"

// Handler ...
type Handler struct {
	authHandler      *authHandler
	spaceshipHandler *spaceshipHandler
}

// SetupHandlers ...
func SetupHandlers(container *container.Container) *Handler {
	authHandler := SetupAuthHandler(container.AuthSvc)
	spaceshipHandler := SetupSpaceshipHandler(container.SpaceshipSvc)
	return &Handler{
		authHandler:      authHandler,
		spaceshipHandler: spaceshipHandler,
	}
}
