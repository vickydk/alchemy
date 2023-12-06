package container

import (
	"alchemy/src/infrastructure/gorm"
	"alchemy/src/shared/config"
	Database "alchemy/src/shared/database"
	authSvc "alchemy/src/usecase/auth"
	spaceshipSvc "alchemy/src/usecase/spaceship"
)

type Container struct {
	Config       *config.Config
	AuthSvc      authSvc.Service
	SpaceshipSvc spaceshipSvc.Service
}

func Setup() *Container {
	// ====== Construct Config
	cfg := config.NewConfig("./resources/config.json")
	db := Database.New(cfg.Database)

	userRepo := gorm.UserSetup(db)
	spaceshipRepo := gorm.SpaceshipSetup(db)

	authSvc := authSvc.NewService(userRepo)
	spaceshipSvc := spaceshipSvc.NewService(spaceshipRepo)
	return &Container{
		Config:       cfg,
		AuthSvc:      authSvc,
		SpaceshipSvc: spaceshipSvc,
	}
}
