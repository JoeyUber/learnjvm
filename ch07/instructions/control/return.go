package control

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type RETURN struct{ base.NoOperandsInstruction }  // Return void from method
type ARETURN struct{ base.NoOperandsInstruction } // Return reference from method
type DRETURN struct{ base.NoOperandsInstruction } // Return double from method
type FRETURN struct{ base.NoOperandsInstruction } // Return float from method
type IRETURN struct{ base.NoOperandsInstruction } // Return int from method
type LRETURN struct{ base.NoOperandsInstruction } // Return long from method

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
	invokerFrame := frame.Thread().TopFrame()
	val := frame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}
