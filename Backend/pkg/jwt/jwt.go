package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT for a valid user
func GenerateToken(userID int, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("5bcb23c98f0f4c438bd92b149f1df6d1a03b97e5cbb476a6fd5b2421e924cd5e"))
}

// ValidateToken parses and validates the JWT token string
func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// Access jwtKey at runtime to ensure AppConfig is initialized
		return []byte("5bcb23c98f0f4c438bd92b149f1df6d1a03b97e5cbb476a6fd5b2421e924cd5e"), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
