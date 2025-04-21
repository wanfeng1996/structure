package structure

import (
	"golang.org/x/exp/constraints"
	"slices"
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

func (it Iter[T]) transition() Iter[T] {
	slices.DeleteFunc(it, func(t *T) bool {
		if t == nil {
			return true
		}
		return false
	})
	return it
}

func (it Iter[T]) Index() []T {
	var res []T
	for i := range it {
		if it[i] != nil {
			res = append(res, *it[i])
		}
	}

	return res
}
