package authapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SumirVats2003/workout-api/models"
)

type credentials struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var payload credentials

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := models.User{
		Name:   payload.Name,
		Email:  payload.Email,
		UserId: "7898",
	}

	password := payload.Password
	fmt.Println(password)

	fmt.Fprintf(w, "user : %+v", user)
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to register route")
}
