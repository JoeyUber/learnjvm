package classfile
 import "encoding/binary" 
 type ClassReader struct{
 	data []byte
 }

 func (self *ClassReader) readUint8() uint8{
 	val := self.data[0]
 	self.data = self.data[1:]
 	return val
 }

 func (self *ClassReader) readUint16() unit16{
 	val := binary.BigEndian.Unit16(self.data)
 	self.data = self.data[2:]
 	return val
 }


 func (self *ClassReader) readUint32() unit32{
 	val := binary.BigEndian.Unit32(self.data)
 	self.data = self.data[4:]
 	return val
 }


func (self *ClassReader) readUint64() unit64{
 	val := binary.BigEndian.Unit64(self.data)
 	self.data = self.data[8:]
 	return val
 }

func (self *ClassReader) readUint16s() []unit16{
 	n :=self.readUint16()
 	s :=make([]unit16,n)
 	for i := rang s{
 		s[i] = self.readUint16()
 	} 
 	return s;
 }

func (self *ClassReader) readBytes(n unit32) []byte{
	val := self.data[:n]
	self.data = self.data[n:]
	return val;
}

