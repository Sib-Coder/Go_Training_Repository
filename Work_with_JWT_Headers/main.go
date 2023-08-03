package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	tocken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTEwNjQ1MDcsIm5pY2siOiJSb2NraSJ9.D9cNWYh4WBP18ZMQl_Q8dVT3WxPzyTd0HYA9Pty8blY"
	dan, _ := extractClaims(tocken)
	fmt.Println(dan)
	des, _ := ExtracNickForJWT(tocken)
	fmt.Println(des)
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := "super-secret-auth-key"
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Printf("Invalid JWT Token")
		return nil, false
	}
}

func ExtracNickForJWT(tokenString string) (string, error) {
	var name string
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		name = fmt.Sprint(claims["nick"])
	}

	if name == "" {
		return "", fmt.Errorf("invalid token payload")
	}
	return name, nil
}
