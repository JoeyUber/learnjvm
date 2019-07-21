package classfile
 import "fmt"
 type ClassFile struct{
 	magic  uint32
 	minorVersion uint16
 	majorVersion unit16
 	constantPool ConstantPool
 	accessFlags unit16
 	thisClass unit16
 	superClass unit16
 	interfaces []unit16
 	fields  []*MemberInfo
 	methods  []*MemberInfo
 	attributes []AttributeInfo
 }

 func Parse(classData []byte) (cf *ClassFile ,err error){
 	//defer func
 	cr := &ClassReader{classData}
 	cf := &ClassFile{}
 	cf.read(cr)
 	return
 }

 func (self *ClassFile) read(reader *ClassReader){
 	self.readAndCheckMagic(reader)
 	self.readAndCheckVersion(reader)
    self.contantPool = readContantPool(reader)
    self.accessFlags = reader.readUnit16()
    self.thisClass = reader.readUnit16()
    self.superClass = reader.readUnit16()
    self.interfaces = reader.readUnit16s()
    self.fields = readMembers(self,self.contantPool)
    self.methods = readMembers(self,self.contantPool)
    self.attributes = readAttributes(self,self.contantPool)
 }

 func (self	*ClassFile) MajorVersion() unit16{
 	return self.majorVersion;
 }

 func (self *ClassFile) ClassName() string{
 	return self.constantPool.getClassName(self.thisClass);
 }

 func (self *ClassFile) SuperClassName() string{
 	if self.superClass > 0{
	 	return self.constantPool.getClassName(self.superClass)
	}
	return ""
 }

 func (self *ClassFile) InterfaceNames() []string{
 	interfaceNames := make([]string,len(self.interfaces))
 	for i ,index := range self.interfaces{
 		interfaceNames[i] = self.constantPool.getClassName(index);
 	}
 	return interfaceNames;
 }

 func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
 	magic :=  reader.readUnit32()
 	if magic != 0xCAFEBABE{
 		panic("java.lang.ClassFormatError : magic ! ")
 	}
 }  

 func (self *ClassFile) readAndCheckVersion(reader *ClassReader){
 	self.minorVersion = reader.readUnit16()
 	self.majorVersion = reader.readUnit16()
 	switch self majorVersion {
 		case 45 :
 			return
 		case 46,47,48,49,50,51,52:
 			if self.minorVersion == 0{
 				return
 			}
 	}
 	panic("java.lang.UnsupportedClassVersionError!")
 }

 