package auth

import (
	ctxSess "alchemy/src/shared/utils/context"
)

type Service interface {
	Login(ctxSess *ctxSess.Context, req *LoginRequest) (res *LoginResponse, err error)
}
