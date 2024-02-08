package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")

type TokenPayload struct {
	Username    string `json:"username"`
	ID          int    `json:"id"`
	Position_id string `json:"position_id"`
	Role_id     int    `json:"role_id"`
	Viewer      bool   `json:"viewer"`
	Areas       []uint `json:"areas"`
}

func GenerateJwtToken(payload TokenPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"payload": payload,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
