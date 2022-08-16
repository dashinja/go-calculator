package main

import (
	// "calculator/math/calculate"
	"calculator/math/precalc"
)



var testInputs = []int{16,6,8,4,5,6,9,8,9,10}

var Ops = precalc.Inputs{InputRange: testInputs, Err: false}

func main() {
  
  Ops.Add(Ops.InputRange)
  Ops.Sub(Ops.InputRange)
  Ops.Mult(Ops.InputRange)
  Ops.Div(Ops.InputRange)
  Ops.Mod(Ops.InputRange)
  // Ops.Add(Ops.InputRange)
  // Ops.Add(Ops.InputRange)
//   calculate.Ops.Calculator("add")
//   calculate.Ops.Calculator("sub")
//   calculate.Ops.Calculator("mult")
//  calculate.Ops.Calculator("div")
//   calculate.Ops.Calculator("sq")
//   calculate.Ops.Calculator("mod")
}
