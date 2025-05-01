package authroutes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRoutes(db *sql.DB) {
	muxRouter := mux.NewRouter()
	db.Ping()

	muxRouter.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome to login route")
	}).Methods("POST")

	muxRouter.HandleFunc("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome to register route")
	}).Methods("POST")

	http.ListenAndServe(":3000", nil)
}
