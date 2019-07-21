package main
import "os"
import "fmt"
import "strings"
import "jvmgo/ch03/classpath"
import "jvmgo/ch03/classfile"
 func main(){
 	cmd := parseCmd()
 	  if cmd.versionFlag {
 	  	 jh := os.Getenv("JAVA_HOME")
 	  	 fmt.Printf("version 0.0.1  java_home is %s \n",jh)
 	  }else if cmd.helpFlag || cmd.class == "" {
 	  	 printUsage()
 	  }else{
 	  	 startJVM(cmd)
 	  }

 }


 func startJVM(cmd *Cmd){
 	cp := classpath.Parse(cmd.XjreOption,cmd.cpOption)

 	fmt.Printf(" classpath:%s class:%s args:%v \n ",cmd.cpOption,cmd.class,cmd.args)

 	className := strings.Replace(cmd.class,".","/",-1)

 	classData,_,err := cp.ReadClass(className)
 	if(err != nil){
 	
 		fmt.Printf("Could not found or load main class %v \n",cmd.class)
 		return
 	} 

 	fmt.Printf(" class data %v \n",classData)

 }