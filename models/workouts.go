package models

type Workout struct {
	WorkoutId  string
	PlanId     string
	ExerciseId string
	Reps       int
	Sets       int
	Weight     float64
}
