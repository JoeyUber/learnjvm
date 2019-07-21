package classfile
type MemberInfo struct{
	cp ConstantPool
	accessFlags uint16
	nameIndex uint16
	descriptionIndex uint16
	attributes []AttributInfo
}


func readMembers(reader *ClassReader,cp ConstantPool) []*MemberInfo{
	memberCount := reader.readUnit16()
	members := make([]*MemberInfo,memberCount)
	for i := range members{
		members[i] = readMember(reader,cp)
	}
	return members
}

func readMember(reader *ClassReader,cp ConstantPool) MemberInfo{
	return &MemberInfo{
		cp : cp,
		accessFlags : reader.readUnit16()
		nameIndex : reader.readUnit16()
		descriptionIndex : reader.readUnit16()
		attributes : readAttributes(reader,cp)
	}
}

func (self *MemberInfo) Name() string{
	self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Description() string{
	self.cp.getUtf8(self.descriptionIndex)
}

