package models

type MuscleGroup int

const (
	Chest MuscleGroup = iota
	Biceps
	Triceps
	Shoulders
	Back
	Legs
)

func (m MuscleGroup) String() string {
	return [...]string{"Chest", "Biceps", "Triceps", "Shoulders", "Back", "Legs"}[m-1]
}

func (m MuscleGroup) Index() int {
	return int(m)
}
