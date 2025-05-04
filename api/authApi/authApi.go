package authapi

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/SumirVats2003/workout-api/models"
)

func Login(db *sql.DB, user models.User, password string) {
	fmt.Println(user)
	fmt.Println(password)
}

func Register(db *sql.DB, user models.User, password string) {
	db.Ping()
	fmt.Println(user)
	fmt.Println(password)
	user.UserId = "" // TODO: implement uuid utility

	res, err := db.Exec(
		"INSERT INTO users (id, created_at, name, email, password) VALUES ($1, $2, $3, $4, $5)",
		user.UserId,
		time.Now(),
		user.Name,
		user.Email,
		password,
	)

	if err != nil {
		log.Println("Insert error:", err)
	}

	fmt.Print(res)
}
