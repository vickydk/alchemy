package http

import (
	"alchemy/src/interface/interceptor"
	"alchemy/src/shared/config"
	"alchemy/src/shared/logger"
	"alchemy/src/shared/utils"
	"alchemy/src/shared/utils/context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

func setupMiddleware(server *echo.Echo, cfg *config.Config) {
	server.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqId := c.Request().Header.Get(echo.HeaderXRequestID)
			if len(reqId) == 0 {
				reqId = utils.GenerateThreadId()
			}

			ctxSess := context.New(logger.GetLogger()).
				SetXRequestID(reqId).
				SetAppName("alchemy.API").
				SetAppVersion("0.1").
				SetPort(cfg.Apps.HttpPort).
				SetSrcIP(c.RealIP()).
				SetURL(c.Request().URL.String()).
				SetMethod(c.Request().Method).
				SetHeader(c.Request().Header)

			ctxSess.Lv1("Incoming Request")

			c.Set(context.AppSession, ctxSess)

			return h(c)
		}
	})
	interceptor := interceptor.New()
	server.Use(interceptor.ValidateAccess())

	server.Validator = &DataValidator{ValidatorData: validator.New()}
	server.Use(middleware.Recover())

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE, echo.HEAD, echo.OPTIONS},
		AllowHeaders: []string{
			"Content-Type", "Origin", "Accept", "Authorization", "Content-Length", "X-Requested-With",
			"OS-Type", "Device-Name", "Device-Serial", "OS-Version", "App-Version",
		},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
}

// DataValidator ...
type DataValidator struct {
	ValidatorData *validator.Validate
}

// Validate ...
func (cv *DataValidator) Validate(i interface{}) error {
	return cv.ValidatorData.Struct(i)
}
