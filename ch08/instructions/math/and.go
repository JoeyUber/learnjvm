package math

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type IAND struct{ base.NoOperandsInstruction }
type LAND struct{ base.NoOperandsInstruction }

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}
