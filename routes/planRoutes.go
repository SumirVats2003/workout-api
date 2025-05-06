package routes

import (
	"github.com/SumirVats2003/workout-api/controllers"
	"github.com/gorilla/mux"
)

func planRoutes(muxRouter *mux.Router, controller *controllers.Controller) {
	planRouter := muxRouter.PathPrefix("/user/{userId}").Subrouter()

	planRouter.HandleFunc("/plan/{planId}", controller.GetPlan).Methods("GET")
	planRouter.HandleFunc("/plan/create", controller.CreatePlan).Methods("POST")
	planRouter.HandleFunc("/plan/{planId}/schedule", controller.SchedulePlan).Methods("POST")
	planRouter.HandleFunc("/plan/{planId}", controller.UpdatePlan).Methods("PUT")
	planRouter.HandleFunc("/plan/{planId}", controller.DeletePlan).Methods("DELETE")
	planRouter.HandleFunc("/all-plans", controller.GetAllPlans).Methods("DELETE")
}
