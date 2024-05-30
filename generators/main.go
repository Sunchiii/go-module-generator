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
		fmt.Println("Usage: generate <module_name>")
		os.Exit(1)
	}

	moduleName := flag.Arg(0)
	projectName, err := generators.GetProjectName()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	generators.GenerateModules(moduleName, projectName)
}
