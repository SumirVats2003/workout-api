package models

type PlanStatus int

const (
	Completed PlanStatus = iota
	Ongoing
)

func (p PlanStatus) String() string {
	return [...]string{"Completed", "Ongoing"}[p-1]
}

func (p PlanStatus) Index() int {
	return int(p)
}
