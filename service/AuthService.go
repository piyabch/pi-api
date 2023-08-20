package service

import (
	"fmt"
	"time"

	"github.com/piyabch/pi-api/model"

	"github.com/dgrijalva/jwt-go"
)

type authCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

const jwtIssuer = "pi"
const jwtSecret = "secret"

// Authenticate API caller with authentication data.
// Return the authentication token.
func Authorize(authData *model.AuthData) (*model.AuthResult, error) {
	var authResult model.AuthResult
	// simulate login
	if authData.Email == "admin@example.com" && authData.Password == "defaultpassword" {
		authResult.Token = generateToken(authData.Email)
		return &authResult, nil
	}
	return nil, nil
}

// Validate token from the API caller.
// Return validation result in bool.
func IsAuthTokenValid(token string) bool {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isvalid := t.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	return jwtToken != nil && jwtToken.Valid && err == nil
}

// Generate JWT token from the given email.
// Return the generated token which valid within a period.
func generateToken(email string) string {
	claims := &authCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    jwtIssuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return ""
	}
	return t
}
