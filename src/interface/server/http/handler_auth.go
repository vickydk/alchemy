package http

import (
	"net/http"

	authSvc "alchemy/src/usecase/auth"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	service authSvc.Service
}

func SetupAuthHandler(service authSvc.Service) *authHandler {
	handler := &authHandler{
		service: service,
	}
	if handler.service == nil {
		panic("service is nil")
	}
	return handler
}

// login godoc
// @Summary  create Token
// @Tags     credential
// @Accept   json
// @Produce  json
// @Param    request  body      auth.LoginRequest  true  "login request"
// @Success  200      {object}  auth.LoginResponse  "Data"
// @Failure  400      {object}  http.HTTPError  "Error message"
// @Router   /login [post]
func (s *authHandler) login(ctx echo.Context) error {
	context := Parse(ctx)
	ctxSess := context.CtxSess

	req := &authSvc.LoginRequest{}
	err := ctx.Bind(req)
	if err != nil {
		return httpError(ctx, http.StatusBadRequest, err)
	}

	resp, err := s.service.Login(ctxSess, req)
	if err != nil {
		return httpError(ctx, http.StatusBadRequest, err)
	}

	return httpOk(ctx, http.StatusCreated, resp)
}
