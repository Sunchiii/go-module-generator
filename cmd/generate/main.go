package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Sunchiii/go-module-generator/generators"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Usage: generate <module_name>||init")
		os.Exit(1)
	}

	moduleName := flag.Arg(0)
	var projectName string

	if moduleName == "init" {
		fmt.Print("Enter Project Name: ")
		fmt.Scan(&projectName)

		// validate project name must not empty and must not contain space
		if projectName == "" {
			log.Fatal("Project name must not empty")
		}

		if le := strings.Split(projectName, " "); len(le) > 1 {
			log.Fatal("Project name must not contain space")
		}

		cmd := exec.Command(fmt.Sprintf("go mod init %s", projectName))
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}

		generators.GenerateInitialStructure()
	} else {
		generators.GenerateModules(moduleName)
	}
}
