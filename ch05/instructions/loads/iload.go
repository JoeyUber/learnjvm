package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type ILOAD struct{ base.Index8Instruction }
type ILOAD_0 struct{ base.NoOperandsInstruction }
type ILOAD_1 struct{ base.NoOperandsInstruction }
type ILOAD_2 struct{ base.NoOperandsInstruction }
type ILOAD_3 struct{ base.NoOperandsInstruction }

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	//fmt.Printf("will push %d to stack \n", val)
	frame.OperandStack().PushInt(val)
}

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, self.Index)
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
