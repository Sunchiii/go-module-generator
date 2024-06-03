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
	CreateLoggers(projectName)
	CreatePagination(projectName)
	CreateAppErrs()
	CreateRoutes()
	CreateFiberRoutes(projectName)
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

func CreateAppErrs() {
	pathFolder := "errs"
	if _, err := os.Stat(pathFolder); os.IsNotExist(err) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	path := "errs/"
	file := path + "errors.go"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		destination, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()

		fmt.Fprintf(destination, "package errs\n\n")
		fmt.Fprintf(destination, "import \"net/http\"\n\n")
		fmt.Fprintf(destination, "type AppError struct {\n")
		fmt.Fprintf(destination, "	Status  int\n")
		fmt.Fprintf(destination, "	Message string\n")
		fmt.Fprintf(destination, "}\n\n")
		fmt.Fprintf(destination, "func (a AppError) Error() string {\n")
		fmt.Fprintf(destination, "	return a.Message\n")
		fmt.Fprintf(destination, "}\n\n")
		fmt.Fprintf(destination, "func NewError(code int, errMsg string) error {\n")
		fmt.Fprintf(destination, "	return AppError{\n")
		fmt.Fprintf(destination, "		Status:  code,\n")
		fmt.Fprintf(destination, "		Message: errMsg,\n")
		fmt.Fprintf(destination, "	}\n")
		fmt.Fprintf(destination, "}\n\n")
		fmt.Fprintf(destination, "func ErrorBadRequest(errorMessage string) error {\n")
		fmt.Fprintf(destination, "	return AppError{\n")
		fmt.Fprintf(destination, "		Status:  http.StatusBadRequest,\n")
		fmt.Fprintf(destination, "		Message: errorMessage,\n")
		fmt.Fprintf(destination, "	}\n")
		fmt.Fprintf(destination, "}\n\n")
		fmt.Fprintf(destination, "func ErrorUnprocessableEntity(errorMessage string) error {\n")
		fmt.Fprintf(destination, "	return AppError{\n")
		fmt.Fprintf(destination, "		Status:  http.StatusUnprocessableEntity,\n")
		fmt.Fprintf(destination, "		Message: errorMessage,\n")
		fmt.Fprintf(destination, "	}\n")
		fmt.Fprintf(destination, "}\n\n")
		fmt.Fprintf(destination, "func ErrorInternalServerError(errorMessage string) error {\n")
		fmt.Fprintf(destination, "	return AppError{\n")
		fmt.Fprintf(destination, "		Status:  http.StatusInternalServerError,\n")
		fmt.Fprintf(destination, "		Message: errorMessage,\n")
		fmt.Fprintf(destination, "	}\n")
		fmt.Fprintf(destination, "}\n")

		fmt.Println("Created AppErrs successfully", file)
	} else {
		fmt.Println("File already exists!", file)
	}
}

func CreateLoggers(projectName string) {
	pathFolder := "logs"
	if _, err := os.Stat(pathFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	path := "logs/"
	file := path + "loggers.go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {
		destination, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()

		// Write the logging package code to the file
		fmt.Fprintf(destination, "package logs\n\n")
		fmt.Fprintf(destination, "import (\n")
		fmt.Fprintf(destination, "    \"fmt\"\n")
		fmt.Fprintf(destination, "    \"go.uber.org/zap\"\n")
		fmt.Fprintf(destination, "    \"go.uber.org/zap/zapcore\"\n")
		fmt.Fprintf(destination, ")\n\n")
		fmt.Fprintf(destination, "var log *zap.Logger\n")
		fmt.Fprintf(destination, "var err error\n\n")
		fmt.Fprintf(destination, "func init() {\n")
		fmt.Fprintf(destination, "    config := zap.NewProductionConfig()\n")
		fmt.Fprintf(destination, "    config.EncoderConfig.TimeKey = \"timestamp\"\n")
		fmt.Fprintf(destination, "    config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder\n")
		fmt.Fprintf(destination, "    config.EncoderConfig.StacktraceKey = \"\"\n")
		fmt.Fprintf(destination, "    log, err = config.Build(zap.AddCallerSkip(1))\n")
		fmt.Fprintf(destination, "    if err != nil {\n")
		fmt.Fprintf(destination, "        fmt.Println(err)\n")
		fmt.Fprintf(destination, "        return\n")
		fmt.Fprintf(destination, "    }\n")
		fmt.Fprintf(destination, "}\n\n")
		fmt.Fprintf(destination, "func Info(message string, fields ...zap.Field) {\n")
		fmt.Fprintf(destination, "    log.Info(message, fields...)\n")
		fmt.Fprintf(destination, "}\n\n")
		fmt.Fprintf(destination, "func Error(message interface{}, fields ...zap.Field) {\n")
		fmt.Fprintf(destination, "    switch v := message.(type) {\n")
		fmt.Fprintf(destination, "    case error:\n")
		fmt.Fprintf(destination, "        log.Error(v.Error(), fields...)\n")
		fmt.Fprintf(destination, "    case string:\n")
		fmt.Fprintf(destination, "        log.Error(v, fields...)\n")
		fmt.Fprintf(destination, "    }\n")
		fmt.Fprintf(destination, "}\n\n")
		fmt.Fprintf(destination, "func Debug(message string, fields ...zap.Field) {\n")
		fmt.Fprintf(destination, "    log.Debug(message, fields...)\n")
		fmt.Fprintf(destination, "}\n")

		fmt.Println("Created Loggers successfully", file)
	} else {
		fmt.Println("File already exists!", file)
		return
	}
}

func CreatePagination(projectName string) {
	pathFolder := "paginates"
	if _, err := os.Stat(pathFolder); os.IsNotExist(err) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	path := pathFolder + "/"
	file := path + "pagination.go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {
		destination, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()

		fmt.Fprintf(destination, "package paginates\n\n")
		fmt.Fprintf(destination, "import (\n")
		fmt.Fprintf(destination, "	\"gorm.io/gorm\"\n")
		fmt.Fprintf(destination, "	\"gorm.io/gorm/clause\"\n")
		fmt.Fprintf(destination, ")\n\n")
		fmt.Fprintf(destination, "type PaginateRequest struct {\n")
		fmt.Fprintf(destination, "	Item int `json:\"item\" validate:\"required\"`\n")
		fmt.Fprintf(destination, "	Page int `json:\"page\" validate:\"required\"`\n")
		fmt.Fprintf(destination, "}\n\n")
		fmt.Fprintf(destination, "type PaginatedResponse struct {\n")
		fmt.Fprintf(destination, "	Data        interface{} `json:\"data\"`\n")
		fmt.Fprintf(destination, "	Total       int         `json:\"total\"`\n")
		fmt.Fprintf(destination, "	PerPage     int         `json:\"per_page\"`\n")
		fmt.Fprintf(destination, "	CurrentPage int         `json:\"current_page\"`\n")
		fmt.Fprintf(destination, "	LastPage    int         `json:\"last_page\"`\n")
		fmt.Fprintf(destination, "}\n\n")
		fmt.Fprintf(destination, "func Paginate(db *gorm.DB, model interface{}, paginate PaginateRequest) (*PaginatedResponse, error) {\n")
		fmt.Fprintf(destination, "	var total int64\n")
		fmt.Fprintf(destination, "	db.Model(model).Count(&total)\n")
		fmt.Fprintf(destination, "	lastPage := (int(total) + paginate.Item - 1) / paginate.Item\n")
		fmt.Fprintf(destination, "	offset := (paginate.Page - 1) * paginate.Item\n")
		fmt.Fprintf(destination, "	result := db.Preload(clause.Associations).Limit(paginate.Item).Offset(offset).Find(model)\n")
		fmt.Fprintf(destination, "	if result.Error != nil {\n")
		fmt.Fprintf(destination, "		return nil, result.Error\n")
		fmt.Fprintf(destination, "	}\n")
		fmt.Fprintf(destination, "	pagination := &PaginatedResponse{\n")
		fmt.Fprintf(destination, "		Total:       int(total),\n")
		fmt.Fprintf(destination, "		PerPage:     paginate.Item,\n")
		fmt.Fprintf(destination, "		CurrentPage: paginate.Page,\n")
		fmt.Fprintf(destination, "		LastPage:    lastPage,\n")
		fmt.Fprintf(destination, "		Data:        model,\n")
		fmt.Fprintf(destination, "	}\n")
		fmt.Fprintf(destination, "	return pagination, nil\n")
		fmt.Fprintf(destination, "}\n")

		fmt.Println("Created Pagination successfully:", file)
	} else {
		fmt.Println("File already exists!", file)
	}
}

func CreateRoutes() {
	pathFolder := "routes"
	if _, err := os.Stat(pathFolder); os.IsNotExist(err) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	path := pathFolder + "/"
	file := path + "routes.go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {
		destination, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()

		fmt.Fprintf(destination, "package routes\n\n")
		fmt.Fprintf(destination, "import \"github.com/gofiber/fiber/v2\"\n\n")
		fmt.Fprintf(destination, "type Routes interface {\n")
		fmt.Fprintf(destination, "	Install(app *fiber.App)\n")
		fmt.Fprintf(destination, "}\n")

		fmt.Println("Created Routes successfully:", file)
	} else {
		fmt.Println("File already exists!", file)
	}
}

func CreateFiberRoutes(projectName string) {
	pathFolder := "routes"
	if _, err := os.Stat(pathFolder); os.IsNotExist(err) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	path := pathFolder + "/"
	file := path + "fiber_routes.go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {
		destination, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()

		fmt.Fprintf(destination, "package routes\n\n")
		fmt.Fprintf(destination, "import (\n")
		fmt.Fprintf(destination, "	\"%s/controllers\"\n", projectName)
		fmt.Fprintf(destination, "	\"github.com/gofiber/fiber/v2\"\n")
		fmt.Fprintf(destination, ")\n\n")
		fmt.Fprintf(destination, "type fiberRoutes struct {\n")
		fmt.Fprintf(destination, "	controller controllers.Controller\n")
		fmt.Fprintf(destination, "}\n\n")
		fmt.Fprintf(destination, "func (r fiberRoutes) Install(app *fiber.App) {\n")
		fmt.Fprintf(destination, "	route := app.Group(\"web/\", func(ctx *fiber.Ctx) error {\n")
		fmt.Fprintf(destination, "		return ctx.Next()\n")
		fmt.Fprintf(destination, "	})\n")
		fmt.Fprintf(destination, "	route.Post(\"hello\", r.controller.StartController)\n")
		fmt.Fprintf(destination, "}\n\n")
		fmt.Fprintf(destination, "func NewWebRoutes(\n")
		fmt.Fprintf(destination, "	controller controllers.Controller,\n")
		fmt.Fprintf(destination, ") Routes {\n")
		fmt.Fprintf(destination, "	return &fiberRoutes{\n")
		fmt.Fprintf(destination, "		controller: controller,\n")
		fmt.Fprintf(destination, "	}\n")
		fmt.Fprintf(destination, "}\n")

		fmt.Println("Created fiber_routes.go successfully:", file)
	} else {
		fmt.Println("File already exists!", file)
	}
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
