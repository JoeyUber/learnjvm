package classpath
 import "os"
 import "path/filepath"
 import "fmt"
 type Classpath struct{
 	bootClassPath Entry
 	extClassPath Entry
 	userClassPath Entry
 }

 func (self *Classpath) String() string {
 	return self.userClassPath.String()
 }

 func (self *Classpath) ReadClass(className string) ([]byte,Entry,error){
 	className = className + ".class"
 	if data,entry,err := self.bootClassPath.readClass(className); err == nil {
 		if err != nil {
 			fmt.Printf("bootClassPath not find this class");
 		}else{
 			fmt.Printf("bootClassPath find  class %s \n" ,className)
 		}
 		return data ,entry,err
 	}
 	if data,entry,err := self.extClassPath.readClass(className); err == nil {
 		if err != nil {
 			fmt.Printf("extClassPath not find this class");
 		}
 		return data ,entry,err
 	}
 	return self.userClassPath.readClass(className)
 }

 func Parse(jreOption string,cpOption string) *Classpath{
 	cp := &Classpath{};
 	cp.parseRootAndExtClasspath(jreOption)
 	cp.parseUserClasspath(cpOption)
 	return cp
 }
 
 func (self *Classpath) parseUserClasspath(cpOption string)  {
 	if cpOption == ""{
 		cpOption = "."
 	}
    self.userClassPath = newEntry(cpOption)
 }

 func (self *Classpath) parseRootAndExtClasspath(jreOption string)  {
 	jreDir := getJreDir( jreOption)
    jreRootPath := filepath.Join( jreDir,"lib","*")
    self.bootClassPath = newWildcardEntry(jreRootPath);

    jreExtPath := filepath.Join(jreDir,"lib","ext","*")
    self.extClassPath = newWildcardEntry(jreExtPath)


 }

 func getJreDir(jreOption string) string {
 	if jreOption != "" && exsit(jreOption){
 		return jreOption
 	}
 	if exsit("./jre"){
 		return "./jre"
 	}
 	if jh := os.Getenv("JAVA_HOME") ; jh != ""{
 		return filepath.Join(jh,"jre")
 	}
 	panic("Can not find jre folder!")
 }

 func exsit(path string) bool{
 	if _,err := os.Stat(path); err != nil{
 		if os.IsNotExist(err){
 			return false
 		}
 	}
 	return true
 }
 