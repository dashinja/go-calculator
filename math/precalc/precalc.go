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

// Sums all elements of nums
func (I Inputs) Add(nums []int) int {
	I.Answer = 0
	for _, v := range nums {
		I.Answer += v
	}

	fmt.Printf("Answer: %d\n", I.Answer)
	return I.Answer
}

// Subtracts all elements of nums[1:] from nums[0]
func (I Inputs) Sub(nums []int) int {
	I.Answer = nums[0]
	for _, v := range nums {
		I.Answer -= v
	}

	fmt.Printf("Answer: %d\n", I.Answer)
	return I.Answer
}

// Multiplies all elements of nums together
func (I Inputs) Mult(nums []int) int {
	I.Answer = 1
	for _, v := range nums {
		I.Answer *= v
	}

	fmt.Printf("Answer: %d\n", I.Answer)
	return I.Answer
}

// Dives nums[0] by nums[1] etc sequentially
func (I Inputs) Div(nums []int) int {
	I.Answer = nums[0]
	for _, v := range nums {
		I.Answer /= v
	}

	fmt.Printf("Answer: %d\n", I.Answer)
	return I.Answer
}

// Returns the remainder of the operation nums[0] / nums[1]
func (I Inputs) Mod(nums []int) int {
	fmt.Println("Only the first two numbers in the slice are used for this operation")

	I.Answer = (nums[0] % nums[1])
	fmt.Printf("And the remainder is : %d\n", I.Answer)

	return I.Answer
}
