package references

import "fmt"
import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

// Invoke instance method; dispatch based on class
type INVOKE_VIRTUAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	invokeMethodRef := currentClass.ConstantPool().GetConstant(self.Index).(*heap.MethodRef)
	invokeMethod := invokeMethodRef.ResolvedMethod()
	if invokeMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	ref := frame.OperandStack().GetRefFromTop(invokeMethod.ArgSlotCount())
	if ref == nil {
		if invokeMethodRef.Name() == "println" {
			_println(frame.OperandStack(), invokeMethodRef.Descriptor())
			return
		}
		panic("java.lang.NullPointerException")
	}
	//如果被调用方法是protected并且被调用方法的类P是当前类C的父类并且C和P不在同一个包内
	//并且this的所属类R不是当前类C而且R不是C的子类
	if invokeMethod.IsProtected() &&
		invokeMethod.Class().IsSuperClassOf(currentClass) &&
		invokeMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		currentClass.IsSuperClassOf(ref.Class()) {
		//!ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), invokeMethod.Name(), invokeMethod.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}

func _println(stack *rtda.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(B)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(I)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()
}
