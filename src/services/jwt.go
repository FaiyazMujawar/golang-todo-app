package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte(os.Getenv("JWT_SECRET"))

func SignToken(payload jwt.MapClaims) (string, error) {
	payload["exp"] = time.Now().Add(time.Hour * 24 * 7)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(key)
}
