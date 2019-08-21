package heap

import "jvmgo/ch08/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

// func (self *MemberRef) newMemberRef *MemberRef(refInfo *classfile.ConstantMemberrefInfo){

// }

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndType()
}

func (self *MemberRef) Name() string {
	return self.name
}

func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
