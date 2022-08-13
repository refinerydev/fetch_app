package app

import (
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/refinerydev/fetch_app/helper"
)

type JwtClaims struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func NewRouter(e *echo.Echo) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := middleware.JWTConfig{
		Claims:     &JwtClaims{},
		SigningKey: []byte(os.Getenv("JWT_KEY_SECRET")),
	}

	e.GET("/secret", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtClaims)
		userId := claims.UserId
		role := claims.Role

		return c.JSON(http.StatusOK, helper.M{"user_id": userId, "role": role})
	}, middleware.JWTWithConfig(config))
}
