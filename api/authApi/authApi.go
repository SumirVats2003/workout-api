package authapi

import (
	"database/sql"
	"fmt"

	"github.com/SumirVats2003/workout-api/models"
)

func Login(db *sql.DB, user models.User, password string) {
	fmt.Println(user)
	fmt.Println(password)
}

func Register(db *sql.DB, user models.User, password string) {
	fmt.Println(user)
	fmt.Println(password)
}
