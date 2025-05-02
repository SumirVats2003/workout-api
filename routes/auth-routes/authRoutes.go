package authroutes

import (
	"database/sql"
	authapi "github.com/SumirVats2003/workout-api/api/authApi"
	"github.com/gorilla/mux"
)

func AuthRoutes(muxRouter *mux.Router, db *sql.DB) {
	db.Ping() // INFO: can be removed when more sophisticated db calls are implemented

	authRouter := muxRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authapi.Login).Methods("POST")
	authRouter.HandleFunc("/register", authapi.Register).Methods("POST")
}
