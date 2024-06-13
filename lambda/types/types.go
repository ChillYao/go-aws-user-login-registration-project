package types

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

func NewUser(registerUser RegisterUser) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), 10)
	if err != nil {
		return User{}, err
	}

	return User{
		Username:     registerUser.Username,
		PasswordHash: string(hashedPassword),
	}, nil
}

func ValidatePassword(hashedPassword, plainTextPasswrd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPasswrd))
	return err == nil
}

func CreateToken(user User) string {
	// the output string represents our JWT
	now := time.Now()                           // get the current time
	validUntil := now.Add(time.Hour * 1).Unix() // the token is valid for 1 hour
	claims := jwt.MapClaims{
		"username": user.Username,
		"expires":  validUntil,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)
	secret := "mysecret"

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return ""
	}
	return tokenString
}
