package mathx

import (
	"golang.org/x/exp/constraints"
)

// Abs returns the absolute value of x.
func Abs[T constraints.Signed | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Sum calculates the sum of the input values.
func Sum[T constraints.Integer | constraints.Float | constraints.Complex](x T, y ...T) T {
	sum := x
	for _, yy := range y {
		sum += yy
	}
	return sum
}
