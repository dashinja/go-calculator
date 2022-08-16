package main

import (
	"calculator/math/precalc"
	"fmt"
)

var testInputs = []int{1,2,8,4,5,6,9,8,9,10}

type Inputs struct{
  testRange []int
  answer int
  err bool
}

func main() {
  inputs := Inputs{testRange: testInputs}
  inputs.calculator("add")
  inputs.calculator("sub")
  inputs.calculator("mult")
  inputs.calculator("div")
  inputs.calculator("sq")
}

func (I *Inputs) calculator(operation string) {
  switch operation {
  case "add":
    I.answer = precalc.Add(I.testRange)
   case "sub": 
   I.answer = precalc.Sub(I.testRange) 
  case "mult":
    I.answer = precalc.Mult(I.testRange)
  case "div":
    I.answer = precalc.Div(I.testRange)
  default:
    I.err = true
    fmt.Println("Invalid operation")
  }
  
 if (!I.err) {
   fmt.Printf("Answer: %d\n", I.answer)
 }
}
