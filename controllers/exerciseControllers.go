package controllers

import (
	"encoding/json"
	"net/http"

	exerciseapi "github.com/SumirVats2003/workout-api/api/exerciseApi"
	"github.com/SumirVats2003/workout-api/models"
	"github.com/SumirVats2003/workout-api/utils"
)

type exerciseData struct {
	Name        string             `json:"name"`
	Category    models.Category    `json:"category"`
	MuscleGroup models.MuscleGroup `json:"muscleGroup"`
}

func parseJSONforExercise(payload *exerciseData, w http.ResponseWriter, r *http.Request) (models.Exercise, error) {
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return models.Exercise{}, err
	}
	exercise := models.Exercise{
		ExerciseId:  utils.GenerateUUID(),
		Name:        payload.Name,
		Category:    payload.Category,
		MuscleGroup: payload.MuscleGroup,
	}

	return exercise, nil
}

func (c *Controller) CreateExercise(w http.ResponseWriter, r *http.Request) {
	// TODO: Resolve route hit bug
	var payload exerciseData
	exercise, err := parseJSONforExercise(&payload, w, r)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result := exerciseapi.CreateExercise(c.DB, exercise)

	if result == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Cannot create exercise",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Exercise created successfully",
	})
}
