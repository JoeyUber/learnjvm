package heap

import "jvmgo/ch08/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool,
	refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *InterfaceMethodRef) ResolveInterfaceMethod() *Method {
	if self.method == nil {
		self.resolveInterfaceMethod()
	}
	return self.method
}

func (self *InterfaceMethodRef) resolveInterfaceMethod() {
	d := self.cp.class
	c := self.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupInterfaceMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}

func lookupInterfaceMethod(class *Class, name, descriptor string) *Method {
	for _, method := range class.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	LookupMethodInInterface(class.interfaces, name, descriptor)
	return nil
}
