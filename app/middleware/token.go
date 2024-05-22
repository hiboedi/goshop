package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = os.Getenv("SECRET")

func CreateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte(secretKey)) // Mengubah secretKey menjadi byte slice
	if err != nil {
		return "", err // Mengembalikan kesalahan yang sesungguhnya
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil // Menyesuaikan dengan tipe kembalian secretKey
	})
	if err != nil {
		// Tangani kesalahan parsing
		return fmt.Errorf("Invalid token: %v", err)
	}
	if !token.Valid {
		return fmt.Errorf("Invalid token")
	}
	return nil
}
