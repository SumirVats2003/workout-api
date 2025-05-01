package authroutes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRoutes(db *sql.DB) {
	muxRouter := mux.NewRouter()
	db.Ping() // INFO: can be removed when more sophisticated db calls are implemented

	authRouter := muxRouter.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome to login route")
	}).Methods("POST")

	authRouter.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome to register route")
	}).Methods("POST")

	http.ListenAndServe(":3000", nil)
}
