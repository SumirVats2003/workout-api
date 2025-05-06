package api

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/SumirVats2003/workout-api/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte(os.Getenv("JWT_KEY"))

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func createJsonWebToken(email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"iss": "workout-api",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Login(db *sql.DB, email string, password string) (string, error) {
	var hashedPassword string

	err := db.QueryRow(
		"SELECT password FROM users WHERE email = $1",
		email,
	).Scan(&hashedPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user not found")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	tokenString, err := createJsonWebToken(email)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Register(db *sql.DB, user models.User, password string) sql.Result {
	passwordHash, err := hashPassword(password)

	if err != nil {
		return nil
	}

	res, err := db.Exec(
		"INSERT INTO users (id, created_at, name, email, password) VALUES ($1, $2, $3, $4, $5)",
		user.UserId,
		time.Now(),
		user.Name,
		user.Email,
		passwordHash,
	)
	if err != nil {
		fmt.Println("Insert error:", err)
		return nil
	}

	return res
}
