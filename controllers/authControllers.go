package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/SumirVats2003/workout-api/api"
	"github.com/SumirVats2003/workout-api/models"
	"github.com/SumirVats2003/workout-api/utils"
)

type credentials struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func parseJSONforAuth(payload *credentials, w http.ResponseWriter, r *http.Request) (credentials, error) {
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return credentials{}, err
	}
	credentials := credentials{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	return credentials, nil
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var payload credentials
	credentials, err := parseJSONforAuth(&payload, w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, err := api.Login(c.DB, credentials.Email, credentials.Password)
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
	credentials, err := parseJSONforAuth(&payload, w, r)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user := models.User{
		UserId: utils.GenerateUUID(),
		Name:   credentials.Name,
		Email:  credentials.Email,
	}

	result := api.Register(c.DB, user, credentials.Password)
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
