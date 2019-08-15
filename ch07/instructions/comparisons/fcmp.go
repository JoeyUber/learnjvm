package comparisons

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type FCMPG struct{ base.NoOperandsInstruction }
type FCMPL struct{ base.NoOperandsInstruction }

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
func (self *FCMPG) Execute(frame *rtda.Frame) { _fcmp(frame, true) }
func (self *FCMPL) Execute(frame *rtda.Frame) { _fcmp(frame, false) }
