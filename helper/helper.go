package helper

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Helper to check if file exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// Helper to check if service and controller already exist in the file
func CheckIfExists(f *os.File, serviceName, controllerName string) bool {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, serviceName+"Service") || strings.Contains(line, controllerName+"Controller") {
			return true
		}
	}
	return false
}

// capitalize utility to capitalize the first letter of the controller name
func Capitalize(str string) string {
	return strings.ToUpper(string(str[0])) + str[1:]
}

// Helper function to append a new controller to an existing fiber_routes.go
func ExtendFiberRoutes(file string, newControllerName string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert content to string for easy manipulation
	data := string(content)

	// Check if controller is already added
	if strings.Contains(data, newControllerName+"Controller") {
		fmt.Println("Controller already exists in fiberRoutes struct:", newControllerName)
		return
	}

	// Add the controller to the fiberRoutes struct
	structInsert := fmt.Sprintf("	%s controllers.%sController\n", newControllerName, Capitalize(newControllerName))
	data = strings.Replace(data, "type fiberRoutes struct {\n", "type fiberRoutes struct {\n"+structInsert, 1)

	// Add the controller to the NewFiberRoutes function
	constructorInsert := fmt.Sprintf("	%s controllers.%sController,\n", newControllerName, Capitalize(newControllerName))
	data = strings.Replace(data, "func NewFiberRoutes(\n", "func NewFiberRoutes(\n"+constructorInsert, 1)

	// Add the controller field assignment to the struct initialization
	initInsert := fmt.Sprintf("		%s: %s,\n", newControllerName, newControllerName)
	data = strings.Replace(data, "return fiberRoutes{\n", "return fiberRoutes{\n"+initInsert, 1)

	// Write the updated content back to fiber_routes.go
	if err := ioutil.WriteFile(file, []byte(data), 0644); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("fiber_routes.go updated with new controller:", newControllerName)
	}
}
