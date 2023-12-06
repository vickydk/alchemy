package server

import (
	"alchemy/src/interface/container"
	Http "alchemy/src/interface/server/http"
)

func StartService(container *container.Container) {
	// start http server
	Http.StartHttpService(container)
}
