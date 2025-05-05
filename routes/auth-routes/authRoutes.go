package authroutes

import (
	"database/sql"

	"github.com/SumirVats2003/workout-api/controllers"
	"github.com/gorilla/mux"
)

func AuthRoutes(muxRouter *mux.Router, db *sql.DB) {
	controller := &controllers.Controller{DB: db}
	authRouter := muxRouter.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", controller.Login).Methods("POST")
	authRouter.HandleFunc("/register", controller.Register).Methods("POST")
}
