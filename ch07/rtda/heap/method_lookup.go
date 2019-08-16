package heap

import "fmt"

func LookupMethodInClass(class *Class, name, descriptor string) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			fmt.Printf("find method : %s ,descriptor : %s \n", method.name, method.descriptor)
			if method.name == name && method.descriptor == descriptor {
				fmt.Printf("got method %s \n", name)
				return method
			}
		}
	}
	return nil
}

func LookupMethodInInterface(ifaces []*Class, name, descriptor string) *Method {
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
			method = LookupMethodInInterface(iface.interfaces, name, descriptor)
		}
	}
	return nil
}
