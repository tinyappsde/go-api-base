package auth

import (
	"errors"
	"fmt"
	"time"
	"tinyapps/api-base/handlers"
	"tinyapps/api-base/models"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	UserId string `json:"uid"`
	jwt.StandardClaims
}

func CheckCredentials(env *handlers.Env, email string, password string) (models.User, error) {
	user, err := models.GetUserByEmail(env, email)
	if err != nil {
		return user, err
	}

	fmt.Println(password)
	fmt.Println(user.PasswordHash.String)

	if CheckPasswordHash(password, user.PasswordHash.String) {
		return user, nil
	}

	return user, errors.New("invalid credentials")
}

func IssueToken(user models.User) string {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		panic(err)
	}

	return tokenString
}

func ValidateAccessToken(env *handlers.Env, token string) (models.User, error) {
	var user models.User
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		// if err == jwt.ErrSignatureInvalid {
		// 	// StatusUnauthorized)
		// 	// return
		// }
		return user, err
	}

	if !tkn.Valid {
		// unauthorized
		return user, err
	}

	return models.GetUser(env, claims.UserId)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
