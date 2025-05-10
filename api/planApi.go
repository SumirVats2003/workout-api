package api

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/SumirVats2003/workout-api/models"
)

func GetPlan(db *sql.DB, userId, planId string) models.PlanResponse {
	var status string
	var createdAt time.Time
	var planResponse models.PlanResponse

	db.QueryRow(
		"SELECT status, created_at FROM plans WHERE user_id = $1 AND id = $2",
		userId,
		planId,
	).Scan(
		&status,
		&createdAt,
	)

	res, err := db.Query(
		"SELECT * FROM workouts WHERE plan_id = $1",
		planId,
	)
	if err != nil {
		fmt.Println("cannot retrieve data")
		return models.PlanResponse{}
	}

	res.Scan()

	return planResponse
}

func CreatePlan(db *sql.DB)   {}
func SchedulePlan(db *sql.DB) {}
func UpdatePlan(db *sql.DB)   {}
func DeletePlan(db *sql.DB)   {}
func GetAllPlans(db *sql.DB)  {}
