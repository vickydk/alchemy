package http

import (
	"net/http"
	"strconv"

	spaceshipSvc "alchemy/src/usecase/spaceship"

	"github.com/labstack/echo/v4"
)

type spaceshipHandler struct {
	service spaceshipSvc.Service
}

func SetupSpaceshipHandler(service spaceshipSvc.Service) *spaceshipHandler {
	handler := &spaceshipHandler{
		service: service,
	}
	if handler.service == nil {
		panic("service is nil")
	}
	return handler
}

// createSpaceship godoc
// @Summary  create spaceship
// @Tags     spaceship
// @Accept   json
// @Produce  json
// @Security BearerAuth
// @Param    request  body      spaceship.CreateSpaceshipRequest  true  "login request"
// @Success  201      {object}  http.HTTPResponse  "success true"
// @Failure  400      {object}  http.HTTPResponse  "success false"
// @Router   /spaceship [post]
func (s *spaceshipHandler) createSpaceship(ctx echo.Context) error {
	context := Parse(ctx)
	ctxSess := context.CtxSess

	req := &spaceshipSvc.CreateSpaceshipRequest{}
	err := ctx.Bind(req)
	if err != nil {
		return httpResp(ctx, http.StatusBadRequest, false)
	}

	err = s.service.CreateSpaceship(ctxSess, req)
	if err != nil {
		return httpResp(ctx, http.StatusBadRequest, false)
	}

	return httpResp(ctx, http.StatusCreated, true)
}

// updateSpaceship godoc
// @Summary  update spaceship
// @Tags     spaceship
// @Accept   json
// @Produce  json
// @Security BearerAuth
// @Param    id  path    string  true "spaceship id"
// @Param    request  body      spaceship.CreateSpaceshipRequest  true  "login request"
// @Success  200      {object}  http.HTTPResponse  "success true"
// @Failure  400      {object}  http.HTTPResponse  "success false"
// @Router   /spaceship/id/{id} [post]
func (s *spaceshipHandler) updateSpaceship(ctx echo.Context) error {
	context := Parse(ctx)
	ctxSess := context.CtxSess

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return httpResp(ctx, http.StatusBadRequest, false)
	}

	req := &spaceshipSvc.CreateSpaceshipRequest{}
	err = ctx.Bind(req)
	if err != nil {
		return httpResp(ctx, http.StatusBadRequest, false)
	}

	err = s.service.UpdateSpaceship(ctxSess, id, req)
	if err != nil {
		return httpResp(ctx, http.StatusBadRequest, false)
	}

	return httpResp(ctx, http.StatusOK, true)
}

// findByID godoc
// @Summary  get spaceship
// @Tags     spaceship
// @Accept   json
// @Produce  json
// @Param    id  path    string  true "spaceship id"
// @Param    request  body      spaceship.CreateSpaceshipRequest  true  "login request"
// @Success  200      {object}  spaceship.SpaceshipResponse  "data"
// @Failure  400      {object}  http.HTTPError  "error message"
// @Router   /spaceship/id/{id} [get]
func (s *spaceshipHandler) findByID(ctx echo.Context) error {
	context := Parse(ctx)
	ctxSess := context.CtxSess

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return httpError(ctx, http.StatusBadRequest, err)
	}

	resp, err := s.service.FindByID(ctxSess, id)
	if err != nil {
		return httpError(ctx, http.StatusBadRequest, err)
	}

	return httpOk(ctx, http.StatusOK, resp)
}

// findSpaceship godoc
// @Summary  get list spaceship
// @Tags     spaceship
// @Accept   json
// @Produce  json
// @Param    request  query      spaceship.FilterSpaceship  true  "login request"
// @Success  200      {object}  spaceship.FindSpaceshipResponse  "data"
// @Failure  400      {object}  http.HTTPError  "error message"
// @Router   /spaceship [get]
func (s *spaceshipHandler) findSpaceship(ctx echo.Context) error {
	context := Parse(ctx)
	ctxSess := context.CtxSess

	req := &spaceshipSvc.FilterSpaceship{}
	err := ctx.Bind(req)
	if err != nil {
		return httpError(ctx, http.StatusBadRequest, err)
	}

	resp, err := s.service.FindSpaceship(ctxSess, req)
	if err != nil {
		return httpError(ctx, http.StatusBadRequest, err)
	}

	return httpOk(ctx, http.StatusOK, resp)
}
