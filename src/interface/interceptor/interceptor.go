package interceptor

import (
	"net/http"
	"strings"

	"alchemy/src/shared/models"
	ctxSess "alchemy/src/shared/utils/context"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Interceptor struct {
}

type claims struct {
	AccountID int64  `json:"account_id"`
	Role      int64  `json:"role"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

const signingSecret = "thisisaverylongbutsecuretokenstring"

func New() *Interceptor {
	return &Interceptor{}
}

func (interceptor *Interceptor) ValidateAccess() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if interceptor.isHealthCheck(c) {
				return next(c)
			}

			if interceptor.isSwagger(c) {
				return next(c)
			}

			if interceptor.skipCheckSession(c) {
				return next(c)
			}

			if interceptor.checkSession(c) {
				return next(c)
			}

			return echo.NewHTTPError(http.StatusUnauthorized)
		}
	}
}

func (interceptor *Interceptor) isSwagger(c echo.Context) bool {
	return strings.HasPrefix(c.Request().URL.String(), "/swagger")
}

func (interceptor *Interceptor) isHealthCheck(c echo.Context) bool {
	if c.Request().URL.String() == "/" || c.Request().URL.String() == "/ping" {
		return true
	}
	return false
}

func (interceptor *Interceptor) skipCheckSession(c echo.Context) bool {
	if strings.HasPrefix(c.Request().URL.String(), "/api/v1/login") ||
		strings.HasPrefix(c.Request().URL.String(), "/api/v1/public") {
		return true
	}
	return false
}

func (interceptor *Interceptor) checkSession(c echo.Context) bool {
	data := c.Get(ctxSess.AppSession)
	ctxSess := data.(*ctxSess.Context)
	authHeader := c.Request().Header.Get("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return false
	} else {
		tokenString := authHeaderParts[1]
		claimsToken := &claims{}
		token, _ := jwt.ParseWithClaims(tokenString, claimsToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(signingSecret), nil
		})

		if token == nil {
			return false
		}

		if token.Valid {
			user := models.AccountSession{
				AccountID: claimsToken.AccountID,
				Email:     claimsToken.Email,
				Name:      claimsToken.Name,
			}
			ctxSess.UserSession = user
			return true
		} else {
			return false
		}
	}
}
