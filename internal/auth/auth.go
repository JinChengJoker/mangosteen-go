package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var (
	key []byte
	t   *jwt.Token
	s   string
)

func NewJWT(uid uint) (string, error) {
	key = []byte(viper.GetString("auth.secretKey"))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": uid,
		})
	s, err := t.SignedString(key)
	return s, err
}
