package main

import (
	"fmt"
	"jvmgo/ch07/classpath"
	"jvmgo/ch07/rtda/heap"
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
	classLoader := heap.NewClassLoader(cp)
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpreter(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
