package exerciseapi

import (
	"database/sql"
	"fmt"

	"github.com/SumirVats2003/workout-api/models"
)

func CreateExercise(db *sql.DB, exercise models.Exercise) sql.Result {
	res, err := db.Exec(
		"INSERT INTO exercises (id, name, category, muscle_group) VALUES ($1, $2, $3, $4)",
		exercise.ExerciseId,
		exercise.Name,
		exercise.Category,
		exercise.MuscleGroup,
	)

	if err != nil {
		fmt.Println("Insert error:", err)
		return nil
	}

	return res
}
