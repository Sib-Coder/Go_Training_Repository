package endpoint

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

var SECRET = []byte("super-secret-auth-key") //вынести в env
var api_key = "1234"                         //вместо вот этого нужна проверка на существование пользователя в базе

func CreateJWT() (string, error) { //бизнес логика

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix() //дата
	claims["nick"] = "sib-coder"                     //добавление ника пользователя в jwt

	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

func GetJwt(c echo.Context) error {
	if c.Request().Header["Access"] != nil {
		if c.Request().Header["Access"][0] != api_key {
			return nil
		} else {
			token, err := CreateJWT()
			if err != nil {
				return err
			}
			return c.String(http.StatusOK, token)
		}
		return nil
	}
	return nil
}

func Home(ctx echo.Context) error {
	s := fmt.Sprintf("super secret area")
	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return errors.New("Error CTX Server")
	}
	return nil

}
