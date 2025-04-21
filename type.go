package structure

import (
	"golang.org/x/exp/constraints"
	"time"
)

type Numerical interface {
	constraints.Integer | constraints.Float
}

// Type 约束
type Type interface {
	Numerical | ~string | time.Time
}
type Iter[T Type] []*T

func transition(tt []*time.Time) TimeSort {
	var ts TimeSort
	for i := range tt {
		if tt[i] == nil {
			continue
		}
		ts = append(ts, *tt[i])
	}

	return ts
}
