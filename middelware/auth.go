package middelware

import (
	"fmt"
	"net/http"
	"placement/models"
	"placement/services"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var (
	jwtServer services.JwtService
)

func init() {
	jwtServer = services.NewJwtService()
}

// IsAccessTokenValid middelware validates JWT token
func IsAccessTokenValid(c *fiber.Ctx) error {
	return validate(c, jwtServer.VerifyAccessToken)
}

// IsRefreshTokenValid middelware validates refresh token
func IsRefreshTokenValid(c *fiber.Ctx) error {
	return validate(c, jwtServer.VerifyRefreshToken)
}

func validate(c *fiber.Ctx, verifier func(t string) (*jwt.Token, error)) error {
	authorization := c.Get("Authorization")
	tokenString := strings.Split(authorization, " ")
	if len(tokenString) < 2 {
		err := models.UnAuthorizeError{
			Status:  http.StatusUnauthorized,
			Message: "Invalid authorization",
		}
		return c.Status(err.Status).JSON(err)
	}
	t := tokenString[1]
	token, err := verifier(t)
	if err != nil {
		err := models.UnAuthorizeError{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		}
		return c.Status(err.Status).JSON(err)
	}
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims["roll"])
		c.Locals("roll", claims["roll"])
		c.Locals("userID", claims["uerID"].(string))
		return c.Next()
	}
	e := models.UnAuthorizeError{
		Status:  http.StatusUnauthorized,
		Message: "Invalid authorization",
	}
	return c.Status(e.Status).JSON(err)
}
