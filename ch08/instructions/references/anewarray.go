package references

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
	"jvmgo/ch08/rtda/heap"
)

type ANEW_ARRAY struct {
	base.Index16Instruction
}

func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	ref := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := ref.ResolvedClass()
	arrClass := class.ArrayClass()
	count := stack.PopInt()
	arrRef := arrClass.NewArray(uint(count))
	stack.PushRef(arrRef)
}
