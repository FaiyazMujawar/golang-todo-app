package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte(os.Getenv("JWT_SECRET"))

const expirationTime = time.Hour * 24 * 7

func SignToken(payload jwt.MapClaims) (string, error) {
	payload["exp"] = time.Now().Add(expirationTime).UnixNano()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(key)
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		// Check if token was signed with same algorithm
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("something went wrong in token validation")
	}

	// Check token expiration
	exp := claims["exp"].(float64)
	if float64(time.Now().UnixNano()) > exp {
		return nil, fmt.Errorf("token expired")
	}
	return claims, nil
}
