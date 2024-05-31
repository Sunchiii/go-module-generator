package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Sunchiii/go-module-generator/generators"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Usage: generate <module_name>||init")
		os.Exit(1)
	}

	moduleName := flag.Arg(0)

	if moduleName == "init" {
		generators.GenerateInitialStructure()
	} else {
		generators.GenerateModules(moduleName)
	}
}
