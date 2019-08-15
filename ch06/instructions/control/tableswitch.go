package control

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type TABLE_SWITH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (self *TABLE_SWITH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetCount := reader.ReadInt32()
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetCount)
}

func (self *TABLE_SWITH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	var offset int
	if i >= self.low && i <= self.high {
		offset = int(self.jumpOffsets[i-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}
