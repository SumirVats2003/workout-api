package routes

import (
	"database/sql"
	"github.com/SumirVats2003/workout-api/dbconnector"
)

func connectPostgres() *sql.DB {
	db := dbconnector.OpenDBConnection()
	return db
}
