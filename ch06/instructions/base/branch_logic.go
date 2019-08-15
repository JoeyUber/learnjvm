package base

import (
	"jvmgo/ch06/rtda"
)

func Branch(frame *rtda.Frame, offset int) {
	if frame == nil {
		panic("frame is nil")
	}
	if frame.Thread() == nil {
		panic("frame.thread is nil")
	}
	pc := frame.Thread().PC()
	//fmt.Printf("current pc %d \n", pc)
	nextPC := pc + offset
	//fmt.Printf("nextpc %d \n", nextPC)
	frame.SetNextPC(nextPC)
}
