package routes

import (
	"net/http"

	"github.com/SumirVats2003/workout-api/controllers"
	"github.com/SumirVats2003/workout-api/dbconnector"
	authRoutes "github.com/SumirVats2003/workout-api/routes/auth-routes"
	exerciseroutes "github.com/SumirVats2003/workout-api/routes/exercise-routes"
	"github.com/gorilla/mux"
)

func InitRoutes() {
	router := mux.NewRouter()
	db := dbconnector.OpenDBConnection()

	controller := &controllers.Controller{DB: db}
	authRoutes.AuthRoutes(router, db, controller)
	exerciseroutes.ExerciseRoutes(router, db, controller)

	http.ListenAndServe(":3000", router)
}
