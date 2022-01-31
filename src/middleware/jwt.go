package middleware

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Interface interface {
	GenerateToken(username string, password string) string
	ValidateToken(token string) (result *jwt.Token, err error)
}

type JWT struct {
}

func NewJWT() *JWT {
	return &JWT{}
}

type authCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func getSecretKey() string {
	// secret := os.Getenv("SECRET")
	// if secret != "" {
	// 	secret = "KK_SECRET"
	// }

	// return secret
	return "KK_SECRET"
}

func (mJWT *JWT) GenerateToken(username string, password string) string {

	claims := &authCustomClaims{
		Name: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encrypToken, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		panic(err)
	}

	return encrypToken
}

func (mJWT *JWT) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(getSecretKey()), nil
	})
}
