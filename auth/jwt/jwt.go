package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	SECRET = "Fuck USA"
)

func DefaultSessionExp () int64 { return time.Now().Add(time.Hour * 24 * 30).Unix() }

func GenerateToken(clms map[string]interface{}) (string, error) {
	/* Create token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Set claims */
	claims := token.Claims.(jwt.MapClaims)
	for key, value := range clms {
		claims[key] = value
	}

	/* Generate encoded token and send it as response */
	return token.SignedString([]byte(SECRET))
}
