package heap

import "jvmgo/ch06/classfile"

type ClassMember struct {
	accessFlag uint16
	name       string
	descriptor string
	class      *Class
}

func (self *ClassMember) copyMemberInfo(mi *classfile.MemberInfo) {
	self.accessFlag = mi.AccessFlags()
	self.name = mi.Name()
	self.descriptor = mi.Description()
}

func (self *ClassMember) Name() string {
	return self.name
}

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}

func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlag&ACC_STATIC
}

func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlag&ACC_FINAL
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlag&ACC_PUBLIC
}

func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlag&ACC_PROTECTED
}

func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlag&ACC_PRIVATE
}
