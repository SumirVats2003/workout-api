package exerciseroutes

import (
	"database/sql"

	"github.com/SumirVats2003/workout-api/controllers"
	"github.com/gorilla/mux"
)

func ExerciseRoutes(muxRouter *mux.Router, db *sql.DB, controller *controllers.Controller) {
	exerciseRouter := muxRouter.PathPrefix("/exercise").Subrouter()
	exerciseRouter.HandleFunc("/get/{name}", controller.GetExercise).Methods("GET")
	exerciseRouter.HandleFunc("/create", controller.CreateExercise).Methods("POST")
}
