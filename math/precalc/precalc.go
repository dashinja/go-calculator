package precalc

var result int

func Add(nums []int) int {
	result = 0
	for _, v := range nums {
		result += v
	}
	return result
}

func Sub(nums []int) int {
	result = nums[0]
	for _, v := range nums {
		result -= v
	}
	return result
}

func Mult(nums []int) int {
	result = 1
	for _, v := range nums {
		result *= v
	}
	return result
}

func Div(nums []int) int {
	result = nums[0]
	for _, v := range nums {
		result /= v
	}
	return result
}