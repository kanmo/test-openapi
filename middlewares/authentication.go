package middlewares

import (
	"database/sql"
	"encoding/base64"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"test-openapi/repository"
)

func AuthHandler(conn *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
			if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing or malformed token")
			}

			cred := strings.TrimPrefix(authHeader, "Basic ")
			decodedCredentials, err := base64.StdEncoding.DecodeString(cred)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid credential")
			}

			credentials := strings.SplitN(string(decodedCredentials), ":", 2)
			if len(credentials) != 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid credential")
			}
			uIDStr := credentials[0]
			pwd := credentials[1]

			ctx := c.Request().Context()
			repo := repository.NewUserRepository(conn)

			uID, err := strconv.Atoi(uIDStr)
			if err != nil {
				log.Printf("invalid uid: %s", uIDStr)
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid user id")
			}

			user, err := repo.FindByID(ctx, uID)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "user not found")
			}

			if err = bcrypt.CompareHashAndPassword([]byte(user.PWHash), []byte(pwd)); err != nil {
				log.Printf("invalid password: %s", pwd)
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
			}

			c.Set("subject", user.Name)

			//if user.PWHash != pwd {
			//	log.Printf("invalid password: %s", pwd)
			//	return echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
			//}

			return next(c)
		}
	}
}
