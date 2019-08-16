package heap

import (
	"fmt"
	"jvmgo/ch07/classfile"
)

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, cpInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&cpInfo.ConstantMemberrefInfo)
	return ref
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.method = self.resolvedMethod()
	}
	return self.method
}

func (self *MethodRef) resolvedMethod() *Method {
	d := self.cp.class
	c := self.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	return method
}

func (self *MemberRef) ResolvedInterfaceMethod() *Method {
	c := self.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	for _, method := range c.methods {
		if method.name == self.name && method.descriptor == self.descriptor {
			return method
		}
	}
	return LookupMethodInInterface(c.interfaces, self.name, self.descriptor)
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	fmt.Printf("lookup method : %s in class : %s \n", name, class.name)
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = LookupMethodInInterface(class.interfaces, name, descriptor)
	}
	return method
}
