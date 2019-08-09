package extended

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type IFNULL struct{ base.BranchInstruction }    // Branch if reference is null
type IFNONNULL struct{ base.BranchInstruction } // Branch if reference not null

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		self.Branch(self.Offset)
	}
}

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		self.Branch(self.Offset)
	}
}
