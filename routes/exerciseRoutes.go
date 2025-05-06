package routes

import (
	"github.com/SumirVats2003/workout-api/controllers"
	"github.com/gorilla/mux"
)

func exerciseRoutes(muxRouter *mux.Router, controller *controllers.Controller) {
	exerciseRouter := muxRouter.PathPrefix("/exercise").Subrouter()
	exerciseRouter.HandleFunc("/get/{name}", controller.GetExercise).Methods("GET")
	exerciseRouter.HandleFunc("/create", controller.CreateExercise).Methods("POST")
}
