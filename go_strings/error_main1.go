package main

import "fmt"

type ErrNegativeSqrt float64

func (s ErrNegativeSqrt) Error() string {
	if s < 0 {
		return fmt.Sprintf("cannot Sqrt negative number %f", s)
	}
	return "xxx"
}

func main() {
	s := ErrNegativeSqrt(-2.0).Error()
	fmt.Println(s)
}
