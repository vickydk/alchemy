package spaceship

import ctxSess "alchemy/src/shared/utils/context"

type Service interface {
	CreateSpaceship(ctxSess *ctxSess.Context, req *CreateSpaceshipRequest) (err error)
	UpdateSpaceship(ctxSess *ctxSess.Context, id int64, req *CreateSpaceshipRequest) (err error)
	FindSpaceship(ctxSess *ctxSess.Context, filter *FilterSpaceship) (resp *FindSpaceshipResponse, err error)
	FindByID(ctxSess *ctxSess.Context, id int64) (resp interface{}, err error)
}
