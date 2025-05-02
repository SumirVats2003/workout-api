package models

type Plan struct {
	planId string
	userId string
	timestamp int
	status PlanStatus
}
