package models

import "time"

type Plan struct {
	planId string
	userId string
	createdAt time.Time
	status PlanStatus
}
