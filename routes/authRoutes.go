package routes

import (
	"fmt"
	"net/http"
)

func AuthRoutes() {
	db := connectPostgres()
	http.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		db.Ping()
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.ListenAndServe(":3000", nil)
}
