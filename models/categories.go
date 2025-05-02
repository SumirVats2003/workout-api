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
