package main

import (
	"fmt"
	// "github.com/dashinja/calculator/math/precalc"
	"calculator/math/precalc"
)

var testRange = []int{1,2,3,4,5,6,7,8,9,10}

func main() {
  calculator()
}

func calculator() {
  fmt.Println("Run Calculator")
  fmt.Printf("Output of %d", precalc.Add(testRange))
}
