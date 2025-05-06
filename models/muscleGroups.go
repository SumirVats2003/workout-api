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

func ParseMuscleGroup(s string) MuscleGroup {
	var muscleGroupMap = map[string]MuscleGroup{
		"Chest":     Chest,
		"Biceps":    Biceps,
		"Triceps":   Triceps,
		"Shoulders": Shoulders,
		"Back":      Back,
		"Legs":      Legs,
	}
	if mg, ok := muscleGroupMap[s]; ok {
		return mg+1
	}
	return -1
}
