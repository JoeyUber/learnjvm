package heap

import "strings"

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	md := &MethodDescriptor{}
	if descriptor[0:2] == "()" {
		md.paramterTypes = make([]string, 0)
		return md
	} else {
		paramstr := descriptor[1:strings.Index(descriptor, ")")]
		md.paramterTypes = strings.Split(paramstr, ";")
	}
	md.returnType = descriptor[strings.Index(descriptor, ")")+1 : len([]rune(descriptor))-1]
	return md
}
