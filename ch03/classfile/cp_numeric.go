package classfile
import "math"

type ConstantIntegerInfo struct{
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	 bytes := reader.readUnit32()
	 self.val = int32(bytes)
}


type ConstantFloatInfo{
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader){
	bytes := reader.readUnit32()
	self.val = math.Float32frombits(bytes)
}

type ConstantLongInfo{
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader){
	bytes := reader.readUnit64()
	self.val = int64(bytes)
}


type ConstantDoubleInfo{
	val float64
}

func (self	*ConstantDoubleInfo) readInfo(reader *ClassReader){
	bytes := reader.readUnit64()
	self.val = math.Float64frombits(bytes)
}

