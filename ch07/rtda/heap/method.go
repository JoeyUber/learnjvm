package heap

import "jvmgo/ch07/classfile"

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, cfFields []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfFields))
	for i, cfField := range cfFields {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfField)
		methods[i].copyAttributes(cfField)
		methods[i].calcArgSlotCount()
	}
	return methods
}

func (self *Method) calcArgSlotCount() {
	parseMethodDescriptor := parseMethodDescriptor(self.descriptor)
	for _, param := range parseMethodDescriptor.paramterTypes {
		self.argSlotCount++
		if param == "D" || param == "J" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++
	}
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if attr := cfMethod.CodeAttribute(); attr != nil {
		self.maxStack = attr.MaxStack()
		self.maxLocals = attr.MaxLocals()
		self.code = attr.Code()
	}
}
func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
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

func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlag&ACC_ABSTRACT
}
