// Package for adding two integer numbers

package exercise

import "golang.org/x/exp/constraints"

// Adds two integer numbers and returns result
// more on [link]: https://www.mathsisfun.com/numbers/addition.html

type Number interface {
	constraints.Integer | constraints.Float
}

func Add[T Number](a, b T) T {
	return a + b
}
