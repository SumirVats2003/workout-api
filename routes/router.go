package routes

import (
	"net/http"

	"github.com/SumirVats2003/workout-api/controllers"
	"github.com/SumirVats2003/workout-api/dbconnector"
	"github.com/gorilla/mux"
)

func InitRoutes() {
	router := mux.NewRouter()
	db := dbconnector.OpenDBConnection()

	controller := &controllers.Controller{DB: db}
	AuthRoutes(router, db, controller)
	ExerciseRoutes(router, db, controller)

	http.ListenAndServe(":3000", router)
}
