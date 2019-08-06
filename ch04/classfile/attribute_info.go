package classfile

/**
对应 attribute_info{
	u2 attribute_name_index
	u4 attribute_length
	   info[attribute_length]
}
attribute_name_index 是一个常量池索引，真正的特性名称存在常量池里
特性信息info 不存在常量池而是紧跟在attribute_length后面
*/

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attrCount := reader.readUint16()
	attrs := make([]AttributeInfo, attrCount)
	for i := range attrs {
		attrs[i] = readAttribute(reader, cp)
	}
	return attrs
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attr := newAttributeInfo(attrName, attrLen, cp)
	attr.readInfo(reader)
	return attr
}

func newAttributeInfo(attrName string, attrLen uint32,
	cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
		//	case "LocalVariableTable":
		//		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
