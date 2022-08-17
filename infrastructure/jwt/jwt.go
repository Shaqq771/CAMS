package jwt

import (
	"backend-nabati/infrastructure/shared/constant"
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func JWTChecking(tokenString string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(constant.ErrInvalidJWTSigningMethod)
		} else if method != jwt.SigningMethodHS256 {
			return nil, errors.New(constant.ErrInvalidJWTSigningMethod)
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		err = errors.New(constant.ErrInvalidJWTToken)
		return
	}

	return
}
