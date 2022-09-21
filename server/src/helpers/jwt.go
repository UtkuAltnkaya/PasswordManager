package helpers

import (
	"errors"
	"github/exmaximum/passwordmanager/src/models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenType struct {
	Id    string `json:"id"`
	Admin bool   `json:"admin"`
	Exp   int64  `json:"exp"`
}

func CreateJwtToken(user *models.User) (string, int64, error) {
	duration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOUR_TIME"))
	exp := time.Now().Add(time.Hour * time.Duration(duration)).Unix()
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}

func VerifyToken(authToken string) (*TokenType, error) {
	if authToken == "" {
		return nil, errors.New("token is not valid")
	}
	tokenString := splitToken(authToken)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		expTime := int64(claims["exp"].(float64))
		if time.Now().Unix() > expTime {
			return nil, errors.New("unauthorized") //Check
		}

		return &TokenType{
			Id:  claims["user_id"].(string),
			Exp: expTime,
		}, nil
	}
	return nil, err
}

func splitToken(token string) string {
	bToken := strings.Split(token, " ")
	if len(bToken) == 2 {
		return bToken[1]
	}
	return token
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
