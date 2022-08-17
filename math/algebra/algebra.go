package algebra

type XInputs struct {
	Answer int
	Err    bool
}

// NOTE: What are you going to do with this again?
type XInputInterface interface {
	algebraicForm(a int, c int) (x int)
}

// a = x + c where x is unknown value int
func (I *XInputs) Addx(a int, c int) (x int) {
	I.Answer = a - c
	return I.Answer
}

// a = x - c where x is unknown value int
func (I *XInputs) Subx(a int, c int) (x int) {
	I.Answer = a + c
	return I.Answer
}

// a = xc where x is unknown value int
func (I *XInputs) Multx(a int, c int) (x int) {
	I.Answer = a / c
	return I.Answer
}

// a = x/c where x is unknown value int
func (I *XInputs) Divx(a int, c int) (x int) {
	I.Answer = a * c
	return I.Answer
}
