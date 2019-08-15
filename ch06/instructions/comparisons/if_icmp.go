package comparisons

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// Branch if int comparison succeeds
type IF_ICMPEQ struct{ base.BranchInstruction }
type IF_ICMPNE struct{ base.BranchInstruction }
type IF_ICMPLT struct{ base.BranchInstruction }
type IF_ICMPLE struct{ base.BranchInstruction }
type IF_ICMPGT struct{ base.BranchInstruction }
type IF_ICMPGE struct{ base.BranchInstruction }

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	if v1 != v2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	if v1 < v2 {
		base.Branch(frame, self.Offset)
	}
}
