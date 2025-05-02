package authapi

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to login route")
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to register route")
}
