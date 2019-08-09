package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type SWAP struct{ base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	slot1 := operandStack.PopSlot()
	slot2 := operandStack.PopSlot()
	operandStack.PushSlot(slot1)
	operandStack.PushSlot(slot2)
}
