package authroutes

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	authapi "github.com/SumirVats2003/workout-api/api/authApi"
	"github.com/SumirVats2003/workout-api/models"
	"github.com/gorilla/mux"
)

type credentials struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func parseJSON(payload *credentials, w http.ResponseWriter, r *http.Request) (models.User, string, error) {
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return models.User{}, "", err
	}
	user := models.User{
		Name:  payload.Name,
		Email: payload.Email,
	}
	password := payload.Password

	return user, password, nil
}

func AuthRoutes(muxRouter *mux.Router, db *sql.DB) {
	authRouter := muxRouter.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var payload credentials
		user, password, err := parseJSON(&payload, w, r)

		if err != nil {
			log.Fatal(err.Error())
			return
		}

		authapi.Login(db, user, password)
	}).Methods("POST")

	authRouter.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		var payload credentials
		user, password, err := parseJSON(&payload, w, r)

		if err != nil {
			log.Fatal(err.Error())
			return
		}

		authapi.Register(db, user, password)
	}).Methods("POST")
}
