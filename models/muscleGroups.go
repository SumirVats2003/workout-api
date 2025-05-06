package models

type MuscleGroup int

const (
	Chest MuscleGroup = iota + 1
	Biceps
	Triceps
	Shoulders
	Abs
	Back
	Legs
)

func (m MuscleGroup) String() string {
	return [...]string{"Chest", "Biceps", "Triceps", "Shoulders", "Abs", "Back", "Legs"}[m-1]
}

func (m MuscleGroup) Index() int {
	return int(m)
}

func ParseMuscleGroup(s string) MuscleGroup {
	var muscleGroupMap = map[string]MuscleGroup{
		"Chest":     Chest,
		"Biceps":    Biceps,
		"Triceps":   Triceps,
		"Shoulders": Shoulders,
		"Abs":       Abs,
		"Back":      Back,
		"Legs":      Legs,
	}
	if mg, ok := muscleGroupMap[s]; ok {
		return mg
	}
	return -1
}
