package main

import (
	"calculator/math/algebra"
	"calculator/math/precalc"
	"fmt"
)

var testInputs = []int{16, 6, 8, 4, 5, 6, 9, 8, 9, 10}

var Ops = precalc.Inputs{InputRange: testInputs, Err: false}

var x = algebra.XInputs{}

func main() {
	Ops.Add(Ops.InputRange)
	Ops.Sub(Ops.InputRange)
	Ops.Mult(Ops.InputRange)
	Ops.Div(Ops.InputRange)
	Ops.Mod(Ops.InputRange)

	x.Addx(10, 3)
	fmt.Println("Result: ", x.Answer)
	x.Multx(10, 3)
	fmt.Println("Result: ", x.Answer)
}
