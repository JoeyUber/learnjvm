package stack

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type DUP struct{ base.NoOperandsInstruction }
type DUP_X1 struct{ base.NoOperandsInstruction }
type DUP_X2 struct{ base.NoOperandsInstruction }
type DUP2 struct{ base.NoOperandsInstruction }
type DUP2_X1 struct{ base.NoOperandsInstruction }
type DUP2_X2 struct{ base.NoOperandsInstruction }

func (self *DUP) Execute(frame *rtda.Frame) {
	slot := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slot)
	frame.OperandStack().PushSlot(slot)
}
