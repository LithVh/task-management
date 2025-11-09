package utils

import (
	"fmt"
	"task-management/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID `json:"userID"`
	jwt.RegisteredClaims
}

func CreateToken(config *config.Config, userID uuid.UUID) (string, error) {
	expTime := time.Now().Add(time.Hour * time.Duration(config.JWT.ExpireHours))

	claims := &Claims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf("CreateToken - %v", err)
	}

	return tokenString, nil
}

func ParseToken(tokenString string, config config.Config) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.JWT.Secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("ParseToken: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token - ParseToken: %v")
	}

	return claims, nil
}
