package authroutes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to login route")
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to register route")
}

func AuthRoutes(muxRouter *mux.Router, db *sql.DB) {
	db.Ping() // INFO: can be removed when more sophisticated db calls are implemented

	authRouter := muxRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", loginUser).Methods("POST")
	authRouter.HandleFunc("/register", registerUser).Methods("POST")
}
