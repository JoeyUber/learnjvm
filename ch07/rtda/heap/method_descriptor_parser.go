package heap

import (
	"strings"
)

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	md := &MethodDescriptor{}
	if descriptor[0:2] == "()" {
		md.paramterTypes = make([]string, 0)
		return md
	} else {
		//fmt.Printf("---------------descriptor %s \n", descriptor)
		paramstr := descriptor[1:strings.Index(descriptor, ")")]
		paramstr = strings.Replace(paramstr, ";", " ", -1)
		md.paramterTypes = strings.Fields(paramstr)

		// fmt.Printf("---------------paramstr %s------------------------- \n", paramstr)
		// fmt.Printf("---------------paramterTypes %v \n", md.paramterTypes)
		// fmt.Printf("---------------paramterTypes count %d \n", len(md.paramterTypes))
		// fmt.Printf("---------------paramterTypes last %s--\n", md.paramterTypes[len(md.paramterTypes)-1])

	}
	md.returnType = descriptor[strings.Index(descriptor, ")")+1 : len([]rune(descriptor))-1]
	return md
}
