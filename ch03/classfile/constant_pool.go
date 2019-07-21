package classfile
type ConstantPool []ConstantInfo
func  readConstantPool(reader *ClassReader) ConstantPool{
		cpCount := reader.readUnit16()
		cp := make([]ConstantInfo,cpCount)
		for 1 = 1; i < cpCount ; i++{
			cp[i] = readConstentInfo(reader,cp)
			switch cp[i].(type){
				case *ConstantLongInfo , *ConstantDoubleInfo:
					i++
			}
		}
	}

func (self *ConstantPool)  getConstantInfo(index unit16) ConstantInfo{
	if cpInfo := self[index]; cpInfo != nil{
		return cpInfo
	}
	panic("Invalid constant pool index!")
}


func (self *ConstantPool)  getNameAndType(index	unit16) (string,string){
	ntInfo := self.getConstantInfo(index).(*ConstantInfoNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.typeIndex)
	return name,_type
}


func (self *ConstantPool)  getClassName(index unit16) string{
	classInfo := self.getConstantInfo(index).(*ConstantInfoClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}


func (self *ConstantPool) getUtf8(index unit16) string{
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}


