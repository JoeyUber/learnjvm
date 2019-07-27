package main

import "os"
import "flag"
import "fmt"

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XjreOption  string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "this is help")
	flag.BoolVar(&cmd.versionFlag, "version", false, "this is version")
	flag.StringVar(&cmd.cpOption, "classpath", "", "this is classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "this is cp")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to Xjre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[:1]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage : %s [-option] class [arg...]\n", os.Args[0])
}
