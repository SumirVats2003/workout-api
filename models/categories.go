package models

type Category int

const (
	Cardio Category = iota
	Strength
	Flexibility
)

func (c Category) String() string {
	return [...]string{"Cardio", "Strength", "Flexibility"}[c-1]
}

func (c Category) Index() int {
	return int(c)
}

func ParseCategory(s string) Category {
	var categoryMap = map[string]Category{
		"Cardio": Cardio,
		"Strength": Strength,
		"Flexibility": Flexibility,
	}
	if c, ok := categoryMap[s]; ok {
		return c+1
	}
	return -1
}
