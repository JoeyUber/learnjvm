package classfile

/**
对应 CONSTANT_NameAndType_info{
	u1 tag
	u2 name_index
	u2 description_index
}
是属性和方法的描述信息
name 是 属性或方法的名称
description 是属性或方法的描述
属性的描述是只属性的类型 如：Ljava/lang/String
方法的描述是指方法入参和返回值的类型 如：(Ljava/lang/Integer;)Ljava/lang/String
*/

type ConstantNameAndTypeInfo struct {
	cp               ConstantPool
	nameIndex        uint16
	descriptionIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptionIndex = reader.readUint16()
}

func (self *ConstantNameAndTypeInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *ConstantNameAndTypeInfo) Type() string {
	return self.cp.getUtf8(self.descriptionIndex)
}
