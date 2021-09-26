package authmw

import (
	"fmt"
	"github.com/ArtemGretsov/golang-server-template/src/config"
	"github.com/golang-jwt/jwt"
)

func CreateJWT(payload JWTPayload) (string, error) {
	serverConfig := config.Get()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": payload.ID,
		"name": payload.Name,
		"login": payload.Login,
	})

	return token.SignedString([]byte(serverConfig["JWT_SECRET"]))
}

func ParseJWT(tokenString string) (JWTPayload, error) {
	serverConfig := config.Get()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return  []byte(serverConfig["JWT_SECRET"]), nil
	})

	if err != nil {
		return JWTPayload{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["id"].(float64))

		return JWTPayload{
			ID: userID,
			Name: claims["name"].(string),
			Login: claims["login"].(string),
		}, nil
	}

	return JWTPayload{}, fmt.Errorf("unexpected JWT payload")
}
