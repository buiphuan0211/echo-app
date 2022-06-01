package util

import (
	"echo-app/config"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	// ID String
	Data map[string]interface{}
	jwt.StandardClaims
}

func GenerateToken(data map[string]interface{}) (string, error) {
	var env = config.GetEnv()

	// create expires time
	expTimeMs, err := strconv.Atoi(env.Jwt.TokenLife)
	if err != nil {
		fmt.Println(err)
	}
	exp := time.Now().Add(time.Millisecond * time.Duration(expTimeMs)).Unix()

	// init claims
	claims := JwtCustomClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	// generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte(env.Jwt.SecretKey))

	return strToken, err
}

func GetJWTPayload(c echo.Context) (map[string]interface{}, error) {
	// get jwt object from context
	fmt.Println(c.Get("user"))

	user := c.Get("user").(*jwt.Token)
	fmt.Println(user.Raw)

	claims := &JwtCustomClaims{}

	// ParseWithClaims
	_, err := jwt.ParseWithClaims(user.Raw, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv().Jwt.SecretKey), nil
	})

	// if err
	if err != nil {
		return nil, err
	}

	return claims.Data, nil
}
