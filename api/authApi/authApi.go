package authapi

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/SumirVats2003/workout-api/models"
	"github.com/SumirVats2003/workout-api/utils"
)

func Login(db *sql.DB, email string, password string) {
	fmt.Println(email)
	fmt.Println(password)
}

func Register(db *sql.DB, user models.User, password string) {
	user.UserId = utils.GenerateUUID()

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
