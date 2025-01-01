package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	SECRET string
}

var defaultconfig = Config{}

// var secretKey = os.Getenv("SECRET_KEY")

func CreateAuthToken(id uint) (string, error) {

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  id,
		"exp":  time.Now().AddDate(0, 3, 0).Unix(),
		"name": "access",
	})

	tokenString, err := token.SignedString([]byte(defaultconfig.SECRET))

	if err != nil {
		return "error", err
	}

	return tokenString, nil
}

func DecodeAuthToken(tokenString string) (uint, error) {
	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g., []byte("my_secret_key")
		return []byte(defaultconfig.SECRET), nil
	})

	if err != nil {
		return 0, err // Return the error
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		// Check the exp
		expirationTime, ok := claims["exp"].(float64)
		if !ok {
			return 0, jwt.ErrTokenInvalidClaims
		}

		if float64(time.Now().Unix()) > expirationTime {
			return 0, jwt.ErrTokenExpired
		}

		// Token is valid, return the subject (sub) claim
		subject, ok := claims["sub"]
		if !ok {
			return 0, jwt.ErrTokenInvalidId
		}
		sub := uint(subject.(float64))

		return sub, nil
	}

	return 0, jwt.ErrTokenSignatureInvalid
}
