package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type IADD struct {
	base.NoOperandsInstruction
}

func (self IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	//fmt.Printf("iadd v1 %d + v2 %d", v1, v2)
	stack.PushInt(v1 + v2)
}
