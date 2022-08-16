package precalc

import (
	"fmt"
)

var result int

type Inputs struct {
	InputRange []int
	Answer     int
	Err        bool
}

func (I Inputs) Add(nums []int) int {
	I.Answer = 0
	for _, v := range nums {
		I.Answer += v
	}

	fmt.Printf("Answer: %d\n", I.Answer)
	return I.Answer
}

func (I Inputs) Sub(nums []int) int {
	I.Answer = nums[0]
	for _, v := range nums {
		I.Answer -= v
	}

	fmt.Printf("Answer: %d\n", I.Answer)
	return I.Answer
}

func (I Inputs) Mult(nums []int) int {
	I.Answer = 1
	for _, v := range nums {
		I.Answer *= v
	}

	fmt.Printf("Answer: %d\n", I.Answer)
	return I.Answer
}

func (I Inputs) Div(nums []int) int {
	I.Answer = nums[0]
	for _, v := range nums {
		I.Answer /= v
	}

	fmt.Printf("Answer: %d\n", I.Answer)
	return I.Answer
}

func (I Inputs) Mod(nums []int) int {
	fmt.Println("Only the first two numbers in the slice are used for this operation")

	I.Answer = (nums[0] % nums[1])
	fmt.Printf("And the remainder is : %d\n", I.Answer)

	return I.Answer
}
