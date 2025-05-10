package models

import "time"

type WorkoutResponse struct {
	Exercise Exercise
	Reps int
	Sets int
	Weight int
}

type PlanResponse struct {
	Workouts []WorkoutResponse
	Status string
	CreatedAt time.Time
}
