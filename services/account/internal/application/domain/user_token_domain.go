package domain

import (
	"fmt"
	"os"
	"time"

	"github.com/expoure/pismo/account/internal/application/constants"
	"github.com/expoure/pismo/account/internal/configuration/customized_errors"
	"github.com/golang-jwt/jwt"
)

func (ud *AccountDomain) GenerateToken() (string, *customized_errors.RestErr) {
	secret := os.Getenv(constants.JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.ID,
		"email": ud.DocumentNumber,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", customized_errors.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token, err=%s", err.Error()))
	}

	return tokenString, nil
}
