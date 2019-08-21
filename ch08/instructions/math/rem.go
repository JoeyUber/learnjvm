package math

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
	"math"
)

type DREM struct{ base.NoOperandsInstruction }
type FREM struct{ base.NoOperandsInstruction }
type IREM struct{ base.NoOperandsInstruction }
type LREM struct{ base.NoOperandsInstruction }

func (sekf *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i1 := stack.PopInt()
	i2 := stack.PopInt()
	if i2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	i := i1 % i2
	stack.PushInt(i)
}

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}
