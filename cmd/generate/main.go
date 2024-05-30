package main

import (
	"fmt"
	"os"

	"github.com/Sunchiii/go-module-generator/generators"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		return
	}
	filename := os.Args[1]
	generators.GenerateModules(filename)
}
