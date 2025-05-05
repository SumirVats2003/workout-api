package exerciseroutes

import (
	"database/sql"

	"github.com/SumirVats2003/workout-api/controllers"
	"github.com/gorilla/mux"
)


func ExerciseRoutes(muxRouter *mux.Router, db *sql.DB) {
	controller := &controllers.Controller{DB: db}
	muxRouter.HandleFunc("/exercises/create", controller.CreateExercise).Methods("POST")
}
