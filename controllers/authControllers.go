package controllers

import (
	"encoding/json"
	"net/http"

	authapi "github.com/SumirVats2003/workout-api/api/authApi"
	"github.com/SumirVats2003/workout-api/models"
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

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var payload credentials
	user, password, err := parseJSON(&payload, w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, err := authapi.Login(c.DB, user.Email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
		var payload credentials
		user, password, err := parseJSON(&payload, w, r)

		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		result := authapi.Register(c.DB, user, password)
		if result == nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "User registration failed",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User registered successfully",
		})
	}
