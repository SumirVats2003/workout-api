package routes

import (
	"github.com/SumirVats2003/workout-api/dbconnector"
	authroutes "github.com/SumirVats2003/workout-api/routes/auth-routes"
)

func InitRoutes() {
	db := dbconnector.OpenDBConnection()
	authroutes.AuthRoutes(db)
}
