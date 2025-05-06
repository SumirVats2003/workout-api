package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) GetPlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	planId := vars["planId"]

	fmt.Println(userId, planId)
}

func (c *Controller) CreatePlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	fmt.Println(userId)
}

func (c *Controller) SchedulePlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	planId := vars["planId"]

	fmt.Println(userId, planId)
}

func (c *Controller) UpdatePlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	planId := vars["planId"]

	fmt.Println(userId, planId)
}

func (c *Controller) DeletePlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	planId := vars["planId"]

	fmt.Println(userId, planId)
}

func (c *Controller) GetAllPlans(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	fmt.Println(userId)
}
