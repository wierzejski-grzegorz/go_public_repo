// Package go_public_repo is a public package created for the purpose of completing the exercise from the book.
// This package contains a function that takes two integers and returns their sum.
package go_public_repo

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes two numbers and returns their sum.
// More information about the function can be found at [AdditionExplanation].
//
// [AdditionExplanation]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](x, y T) T {
	return x + y
}
