package base

import (
	"jvmgo/ch08/rtda"
	"jvmgo/ch08/rtda/heap"
)

/*
根据当前帧获取线程然后在这个线程上为method创建一个新的栈帧
根据method参数个数将当前帧的操作数栈里的值弹出，依次付给新栈帧的本地变量表
*/
func InvokeMethod(invokeFrame *rtda.Frame, method *heap.Method) {
	thread := invokeFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokeFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}
