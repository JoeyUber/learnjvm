package main

import (
	"fmt"
	"jvmgo/ch05/classfile"
	"jvmgo/ch05/classpath"
	"os"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		jh := os.Getenv("JAVA_HOME")
		fmt.Printf("version 0.0.1  java_home is %s \n", jh)
	} else {
		startJVM(cmd)
	}

}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath: %s \n", cp.String())
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpreter(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, method := range cf.Methods() {
		if method.Name() == "main" && method.Description() == "([Ljava/lang/String;)V" {
			return method
		}
	}
	return nil
}
