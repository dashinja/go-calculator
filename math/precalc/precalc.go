package precalc

import (
	"fmt"
)

var result int

func precalc() {
	fmt.Println("I'm precalc")
}

func Add(nums []int) int {
	result = 0
	for _, v := range nums {
		result += v
	}
	return result
}
