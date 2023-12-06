package session

import (
	"time"

	domainUser "alchemy/src/domain/user"
	"github.com/dgrijalva/jwt-go"
)

const signingSecret = "thisisaverylongbutsecuretokenstring"

type claims struct {
	AccountID int64  `json:"account_id"`
	Role      int64  `json:"role"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	jwt.StandardClaims
}
type Session struct{}

func NewBearerToken(user *domainUser.Entity) (string, time.Time) {
	expiry := time.Now().Add(time.Hour * 24 * 365).Unix()
	claims := &claims{
		AccountID: user.ID,
		Name:      user.Name,
		Email:     user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString([]byte(signingSecret)); err == nil {
		return tokenString, time.Unix(expiry, 0)
	}
	return "", time.Unix(expiry, 0)
}

func RefreshToken(user *domainUser.Entity) (string, time.Time) {
	expiry := time.Now().Add(time.Hour * 24 * 370).Unix()
	claims := &claims{
		AccountID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString([]byte(signingSecret)); err == nil {
		return tokenString, time.Unix(expiry, 0)
	}
	return "", time.Unix(expiry, 0)
}
