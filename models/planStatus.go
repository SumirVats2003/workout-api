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

func ParseStatus(s string) PlanStatus {
	var planStatusMap = map[string]PlanStatus{
		"Completed": Completed,
		"Ongoing":   Ongoing,
	}
	if ps, ok := planStatusMap[s]; ok {
		return ps+1
	}
	return -1
}
