package classfile

/**
  对应CONSTANT_String_info {
	u1 tag
	u2 string_index
  }
  string_index 是一个索引 指向常量池的元素位置 最终还是根据常量池的位置读取utf8数据
*/

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
