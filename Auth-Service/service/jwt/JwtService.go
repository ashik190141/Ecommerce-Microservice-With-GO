package jwt

import (
	"Auth-Service/config"
	"Auth-Service/interfaces"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtServices struct{}

var secretKey = []byte(config.LoadEnvData().SecretKey)

func (s *JwtServices) GenerateToken(email string) (interfaces.JWT, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return interfaces.JWT{}, err
	}
	return interfaces.JWT{Token: tokenString}, nil
}

func (s *JwtServices) ValidateToken(token string) (bool, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false, err
	}

	if !parsedToken.Valid {
		return false, fmt.Errorf("invalid token")
	}

	return true, nil
}
