package types

type Corridor struct {
	ID           int
	Length       int
	TraverseTime float64
	Name         string
	Type         string
	Connects     []int
}
