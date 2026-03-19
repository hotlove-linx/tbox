package utils

import (
	"errors"
	"time"

	"admin-backend/config"

	"github.com/golang-jwt/jwt/v5"
)

type AdminClaims struct {
	AdminID uint64 `json:"admin_id"`
	Email   string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateAdminToken(adminID uint64, email string, remember bool) (string, time.Duration, error) {
	cfg := config.GetConfig()

	var expireDuration time.Duration
	if remember {
		expireDuration = time.Duration(cfg.JWT.RememberExpireDays) * 24 * time.Hour
	} else {
		expireDuration = time.Duration(cfg.JWT.ExpireHours) * time.Hour
	}

	claims := AdminClaims{
		AdminID: adminID,
		Email:   email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "admin-backend",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.JWT.Secret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expireDuration, nil
}

func ParseAdminToken(tokenString string) (*AdminClaims, error) {
	cfg := config.GetConfig()

	token, err := jwt.ParseWithClaims(tokenString, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AdminClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.Issuer != "admin-backend" {
		return nil, errors.New("invalid token issuer")
	}

	return claims, nil
}
