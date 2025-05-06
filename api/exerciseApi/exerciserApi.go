package exerciseapi

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/SumirVats2003/workout-api/models"
)

func GetExercise(db *sql.DB, exerciseName string) (models.Exercise, error) {
	var exercise models.Exercise
	var id, name, category, muscleGroup string
	query := "SELECT id, name, category, muscle_group FROM exercises WHERE name = $1"
	err := db.QueryRow(query, exerciseName).Scan(
		&id,
		&name,
		&category,
		&muscleGroup,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Exercise{}, errors.New("exercise not found")
		}
		return models.Exercise{}, err
	}

	exercise.ExerciseId = id
	exercise.Name = name
	exercise.MuscleGroup = models.ParseMuscleGroup(muscleGroup)
	exercise.Category = models.ParseCategory(category)

	return exercise, nil
}

func CreateExercise(db *sql.DB, exercise models.Exercise) sql.Result {
	res, err := db.Exec(
		"INSERT INTO exercises (id, name, category, muscle_group) VALUES ($1, $2, $3, $4)",
		exercise.ExerciseId,
		exercise.Name,
		exercise.Category.String(),
		exercise.MuscleGroup.String(),
	)

	if err != nil {
		fmt.Println("Insert error:", err)
		return nil
	}

	return res
}
