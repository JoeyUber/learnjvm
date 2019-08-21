package heap

import "strings"

type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

/*

 */
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.loadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}

func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.getPackageName() == other.getPackageName()
}

func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) GetPackageName() string {
	return self.getPackageName()
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Class) IsSuperClassOf(sub *Class) bool {
	//TODO 待优化 如果self 是Object 应该直接返回true
	if sub.superClass == nil {
		return false
	}
	if sub.superClass == self {
		return false
	}
	return self.IsSuperClassOf(sub.superClass)
}
