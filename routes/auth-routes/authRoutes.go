package authroutes

import (
	"database/sql"
	"fmt"
	"net/http"
)

func AuthRoutes(db *sql.DB) {
	http.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		db.Ping()
		fmt.Fprintf(w, "welcome to login route")
	})

	http.HandleFunc("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		db.Ping()
		fmt.Fprintf(w, "welcome to register route")
	})

	http.ListenAndServe(":3000", nil)
}
