package models

import "time"

type Plan struct {
	PlanId    string
	UserId    string
	CreatedAt time.Time
	Status    PlanStatus
}
