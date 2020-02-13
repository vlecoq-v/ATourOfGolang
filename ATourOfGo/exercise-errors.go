// Copy your Sqrt function from the earlier exercise and modify it to return an error value.
// Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.
// Create a new type
// type ErrNegativeSqrt float64
// and make it an error by giving it a
// func (e ErrNegativeSqrt) Error() string
// method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".
// Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?
// Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// return "asdf"
	return fmt.Sprintf("%v is a negative number, which are not supported", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		MyError := ErrNegativeSqrt(x)
		return x, MyError
	}
	var prev float64
	z := float64(1)
	for math.Abs(z-prev) > 0.01 {
		prev = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
