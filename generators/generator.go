package generators

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	WORKDIR = "src/"
)

func GenerateInitialStructure() {
	projectName, err := getProjectName()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	CreateConfigEnv(projectName)
}

func CreateConfigEnv(projectName string) {
	pathFolder := "config"
	if _, err := os.Stat(pathFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	path := "config/"
	file := path + "env.go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {
		destination, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()
		fmt.Fprintf(destination, "import (")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	github.com/spf13/viper")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	%s/logs", projectName)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	strings")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, ")")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "func Init() {")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	viper.SetConfigName(\"config\")")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	viper.SetConfigType(\"yaml\")")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	viper.SetConfigPath(\"./\")")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	viper.AutomaticEnv()")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	viper.SetEnvKeyReplacer(strings.NewReplacer(\".\", \"_\"))")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	err := viper.ReadInConfig()")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	if err != nil {")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	fmt.Println(\"ERROR_READING_CONFIG_FILE\", err)")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	logs.Error(\"ERROR_READING_CONFIG_FILE\")")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	return")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "}")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	fmt.Println(\"SUCCESS_READING_CONFIG_FILE\")")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "}")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "func GetEnv(key, defaultValue string) string {")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	readValue := viper.GetString(key)")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	if readValue == \"\" {")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "		return defaultValue")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	}")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	return readValue")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "}")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "func Env(key string) string {")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	readValue := viper.GetString(key)")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "	return readValue")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "}")
	} else {
		fmt.Println("File already exists!", file)
		return
	}

	fmt.Println("Created Config successfully", file)
}

func GenerateModules(filename string) {
	filename = strings.ToLower(filename)

	projectName, err := getProjectName()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	CreateRequests(filename)
	CreateResponses(filename)
	CreateModels(filename)
	CreateRepositories(filename, projectName)
	CreateServices(filename, projectName)
	CreateControllers(filename, projectName)
}

func CreateRequests(filename string) {
	pathFolder := WORKDIR + "requests"
	if _, err := os.Stat(pathFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	path := WORKDIR + "requests/"
	file := path + filename + "_request" + ".go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {

		destination, err := os.Create(file)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()
		fmt.Fprintf(destination, "package requests")
		//fmt.Fprintf(destination, " %s\n", filename)

	} else {
		fmt.Println("File already exists!", file)
		return
	}

	fmt.Println("Created Request successfully", file)
}

func CreateResponses(filename string) {
	pathFolder := WORKDIR + "responses"
	if _, err := os.Stat(pathFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
	path := WORKDIR + "responses/"
	file := path + filename + "_response" + ".go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {

		destination, err := os.Create(file)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()
		fmt.Fprintf(destination, "package responses")

	} else {
		fmt.Println("File already exists!", file)
		return
	}

	fmt.Println("Created Response successfully", file)
}

func CreateModels(filename string) {
	pathFolder := WORKDIR + "models"
	if _, err := os.Stat(pathFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	path := WORKDIR + "models/"
	file := path + filename + ".go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {
		destination, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()
		upperString := strings.Replace(cases.Title(language.Und, cases.NoLower).String(strings.Replace(filename, "_", " ", -1)), " ", "", -1)

		fmt.Fprintf(destination, "package models")
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `import "gorm.io/gorm"`)
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `type %s struct {`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `gorm.Model`)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "}")
	} else {
		fmt.Println("File already exists!", file)
		return
	}

	fmt.Println("Created Model successfully", file)
}

func CreateRepositories(filename string, projectName string) {
	pathFolder := WORKDIR + "repositories"
	if _, err := os.Stat(pathFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	path := WORKDIR + "repositories/"
	file := path + filename + "_repository" + ".go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {
		destination, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()
		upperString := strings.Replace(cases.Title(language.Und, cases.NoLower).String(strings.Replace(filename, "_", " ", -1)), " ", "", -1)
		lowerString := strings.ToLower(string(upperString[0])) + string(upperString[1:len(upperString)])
		//pwd, err := os.Getwd()
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//arr := strings.Split(pwd, "/")
		//projectName := arr[len(arr)-1]

		fmt.Fprintf(destination, "package repositories")
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `import (`)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `"gorm.io/gorm"`)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `"%s/%smodels"`, projectName, WORKDIR)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, ")")
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `type %sRepository interface{`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `//Insert your function interface`)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `type %sRepository struct {db *gorm.DB}`, lowerString)
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `func New%sRepository(db *gorm.DB) %sRepository {`, upperString, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `// db.Migrator().DropTable(models.%s{})`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `db.AutoMigrate(models.%s{})`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `	return &%sRepository{db: db}`, lowerString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)
	} else {
		fmt.Println("File already exists!", file)
		return
	}

	fmt.Println("Created Repository successfully", file)
}

func CreateServices(filename string, projectName string) {
	pathFolder := WORKDIR + "services"
	if _, err := os.Stat(pathFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	path := WORKDIR + "services/"
	file := path + filename + "_service" + ".go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {
		destination, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()
		upperString := strings.Replace(cases.Title(language.Und, cases.NoLower).String(strings.Replace(filename, "_", " ", -1)), " ", "", -1)
		lowerString := strings.ToLower(string(upperString[0])) + string(upperString[1:len(upperString)])
		//pwd, err := os.Getwd()
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//arr := strings.Split(pwd, "/")
		//projectName := arr[len(arr)-1]

		fmt.Fprintf(destination, "package services")
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `import (`)

		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `"%s/%srepositories"`, projectName, WORKDIR)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `)`)

		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `type %sService interface{`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `//Insert your function interface`)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `type %sService struct {`, lowerString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `repository%s repositories.%sRepository`, upperString, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)

		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `func New%sService(`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `repository%s repositories.%sRepository,`, upperString, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "//repo")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, ") %sService {", upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `	return &%sService{`, lowerString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `repository%s :repository%s,`, upperString, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "//repo")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)

		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)
	} else {
		fmt.Println("File already exists!", file)
		return
	}

	fmt.Println("Created Service successfully", file)
}

func CreateControllers(filename string, projectName string) {
	pathFolder := WORKDIR + "controllers"
	if _, err := os.Stat(pathFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	path := WORKDIR + "controllers/"
	file := path + filename + "_controller" + ".go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {
		destination, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()
		upperString := strings.Replace(cases.Title(language.Und, cases.NoLower).String(strings.Replace(filename, "_", " ", -1)), " ", "", -1)
		lowerString := strings.ToLower(string(upperString[0])) + string(upperString[1:len(upperString)])
		//pwd, err := os.Getwd()
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//arr := strings.Split(pwd, "/")
		//projectName := arr[len(arr)-1]

		fmt.Fprintf(destination, "package controllers")
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `import (`)

		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `"%s/%sservices"`, projectName, WORKDIR)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `)`)

		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `type %sController interface{`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `//Insert your function interface`)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `type %sController struct {`, lowerString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `service%s services.%sService`, upperString, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)

		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `func New%sController(`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `service%s services.%sService,`, upperString, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "//services")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, ") %sController {", upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `	return &%sController{`, lowerString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `service%s :service%s,`, upperString, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "//services")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)

		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)
	} else {
		fmt.Println("File already exists!", file)
		return
	}

	fmt.Println("Created Controller successfully", file)
}

// getProjectName reads the project name from the go.mod file
func getProjectName() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				return parts[1], nil
			}
		}
	}

	return "", errors.New("could not determine module name")
}
