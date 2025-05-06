package routes

import (
	"github.com/SumirVats2003/workout-api/controllers"
	"github.com/gorilla/mux"
)

func authRoutes(muxRouter *mux.Router, controller *controllers.Controller) {
	authRouter := muxRouter.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", controller.Login).Methods("POST")
	authRouter.HandleFunc("/register", controller.Register).Methods("POST")
}
