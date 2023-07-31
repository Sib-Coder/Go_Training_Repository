package midleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")
		if strings.EqualFold(val, roleAdmin) {
			log.Println("red button user detected")
		}
		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}

var SECRET = []byte("super-secret-auth-key") //вынести в env

func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		if ctx.Request().Header["Token"] != nil {
			token, err := jwt.Parse(ctx.Request().Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					ctx.String(http.StatusUnauthorized, "not authorized: ")
				}
				return SECRET, nil
			})

			if err != nil {
				ctx.String(http.StatusUnauthorized, "not authorized: "+err.Error())
			}

			if token.Valid {
				next(ctx)
			}
		} else {
			ctx.String(http.StatusUnauthorized, "not authorized: ")
		}
		return nil
	}
}
