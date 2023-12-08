package types

import (
	"time"
)

type Node struct {
	ID        int
	Name      string
	Opens     time.Time
	Closes    time.Time
	ClimbTime float64
	Building  *string
	Connects  []int
}
