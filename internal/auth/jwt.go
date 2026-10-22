package auth

import (
	"errors"
	"rcon-scheduler-panel/internal/errs"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrJWTExpired         = errs.ErrUnauthorized.WithMessage("JWT expired")
	ErrJWTSignatureFailed = errs.ErrUnauthorized.WithMessage("JWT signature is invalid")
)

type Claims struct {
	jwt.RegisteredClaims
}

func GenerateToken(userID int64, secret string) (string, error) {
	now := time.Now()

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatInt(userID, 10),
			ExpiresAt: jwt.NewNumericDate(now.Add(30 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenString, secret string) (int64, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (any, error) {
			return []byte(secret), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithExpirationRequired(),
	)

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return 0, ErrJWTExpired.WithErr(err)
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return 0, ErrJWTSignatureFailed.WithErr(err)
		default:
			return 0, errs.ErrUnauthorized.WithErr(err)
		}
	}

	claims := token.Claims.(*Claims)

	userID, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil || userID <= 0 {
		return 0, errs.ErrUnauthorized
	}

	return userID, nil
}
