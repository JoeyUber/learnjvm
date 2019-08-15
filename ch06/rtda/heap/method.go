package heap

import "jvmgo/ch06/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
	class     *Class
}

func newMethods(class *Class, cfFields []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfFields))
	for i, cfField := range cfFields {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfField)
		methods[i].copyAttributes(cfField)
	}
	return methods
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if attr := cfMethod.CodeAttribute(); attr != nil {
		self.maxStack = attr.MaxStack()
		self.maxLocals = attr.MaxLocals()
		self.code = attr.Code()
	}
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) MaxLocals() uint {
	return self.maxLocals
}

func (self *Method) Class() *Class {
	return self.class
}

func (self *Method) Code() []byte {
	return self.code
}
