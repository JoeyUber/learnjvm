package references

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
)

/*
只能在 私有方法，构造函数，以及父类方法时调用此指令
*/
type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	//获取要调用的类、方法信息
	currentClass := frame.Method().Class()
	invokeMethodRef := currentClass.ConstantPool().GetConstant(self.Index).(*heap.MethodRef)
	invokeMethod := invokeMethodRef.ResolvedMethod()
	invokeMethodClass := invokeMethod.Class()
	/*判断方法在当前情景是否可调用*/
	//判断该方法是否是构造函数
	if invokeMethod.Name() == "<init>" && currentClass != invokeMethodClass {
		panic("java.lang.NoSuchMethodError")
	}
	//判断方法是否是静态
	if invokeMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(invokeMethod.ArgSlotCount())
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	//方法如果是protected则需要判断当前类是不是和被调用方法所属类同包或是同继承于那个类或者
	if invokeMethod.IsProtected() {
		if //currentClass != invokeMethodClass &&
		invokeMethodClass.GetPackageName() != currentClass.GetPackageName() &&
			!invokeMethodClass.IsSuperClassOf(currentClass) &&
			!invokeMethodClass.IsSuperClassOf(ref.Class()) {
			panic("java.lang.IllegalAccessError")
		}
	}

	methodToBeInvoked := invokeMethod
	if currentClass.HasAccSuper() &&
		invokeMethodClass.IsSuperClassOf(currentClass) &&
		invokeMethod.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
			invokeMethodRef.Name(), invokeMethodRef.Descriptor())
	}
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)

}
