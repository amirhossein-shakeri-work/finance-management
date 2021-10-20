package jwt

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kamva/mgm/v3"
	"github.com/sizata-siege/finance-management/user"
)

const (
	SECRET = "Fuck USA"
	CookieName = "jwt_token"
)

/* jwtCustomClaims are custom claims extending default ones. */
// See https://github.com/golang-jwt/jwt for more examples
type CustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}


func DefaultSessionExpUnix () int64 { return time.Now().Add(time.Hour * 24 * 30).Unix() }
func DefaultSessionExp () time.Time { return time.Now().Add(time.Hour * 24 * 30) }

func GenerateToken(mapClaims map[string]interface{}) (string, error) {
	/* Create token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Set claims */
	claims := token.Claims.(jwt.MapClaims)
	for key, value := range mapClaims {
		claims[key] = value
	}

	/* Generate encoded token and send it as response */
	return token.SignedString([]byte(SECRET))
}

type JWT struct {
	Ctx *fiber.Ctx
	Claims *CustomClaims
	User *user.User
	// ID string
	// user (*JWT) *user.User
}

func New (c *fiber.Ctx) *JWT {
	j := &JWT{ Ctx: c }
	j.Claims = j.parseClaims()
	j.User = j.user()
	return j
}

func (j *JWT) parseClaims () *CustomClaims {
	usr := j.Ctx.Locals("user").(*jwt.Token)
	return usr.Claims.(*CustomClaims)
}

func (j *JWT) user () *user.User {
	/* Parse user */
	u := &user.User{}
	if err := mgm.Coll(u).FindByID(j.parseClaims().ID, u); err != nil {
		return nil
	}
	return u
}
