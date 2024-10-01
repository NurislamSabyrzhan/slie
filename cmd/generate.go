package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "g <name>",
	Short: "Generate router, controller, and usecase files",
	Long:  `Generate router, controller, and usecase files with the given name.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		crud, _ := cmd.Flags().GetBool("crud")
		generateFiles(name, crud)
	},
}

var repositoryCmd = &cobra.Command{
	Use:   "r <name>",
	Short: "Generate repository file",
	Long:  `Generate repository file with the given name.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		generateRepository(name)
	},
}

func init() {
	generateCmd.Flags().BoolP("crud", "c", false, "Generate CRUD operations")
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(repositoryCmd)
}

func generateFiles(name string, crud bool) {
	if crud {
		generateRouter(name)
		generateController(name)
		generateUsecase(name)
	} else {
		generateRouterWithoutCRUD(name)
		generateControllerWithoutCRUD(name)
		generateUsecaseWithoutCRUD(name)
	}
}

func generateRouter(name string) {
	tmpl := `package router

import (
	"github.com/gofiber/fiber/v2"
	"newserver/internal/controller"
	"newserver/internal/usecase"
)

func New{{.Name}}Router(app *fiber.App, useCase usecase.{{.Name}}UseCase) {
	ctrl := controller.New{{.Name}}Controller(useCase)
	{{.LowerName}} := app.Group("/{{.LowerName}}")

	{{.LowerName}}.Get("/", ctrl.Get)
	{{.LowerName}}.Post("/", ctrl.Post)
	{{.LowerName}}.Put("/:id", ctrl.Put)
	{{.LowerName}}.Delete("/:id", ctrl.Delete)
}
`
	generateFile("router", name, tmpl)
}

func generateController(name string) {
	tmpl := `package controller

import (
	"github.com/gofiber/fiber/v2"
	"newserver/internal/usecase"
)

type {{.Name}}Controller struct {
	useCase usecase.{{.Name}}UseCase
}

func New{{.Name}}Controller(useCase usecase.{{.Name}}UseCase) *{{.Name}}Controller {
	return &{{.Name}}Controller{useCase: useCase}
}

func (c *{{.Name}}Controller) Get(ctx *fiber.Ctx) error {
	// Implement your logic here
	return ctx.SendString("Get {{.Name}}")
}

func (c *{{.Name}}Controller) Post(ctx *fiber.Ctx) error {
	// Implement your logic here
	return ctx.SendString("Post {{.Name}}")
}

func (c *{{.Name}}Controller) Put(ctx *fiber.Ctx) error {
	// Implement your logic here
	return ctx.SendString("Put {{.Name}}")
}

func (c *{{.Name}}Controller) Delete(ctx *fiber.Ctx) error {
	// Implement your logic here
	return ctx.SendString("Delete {{.Name}}")
}
`
	generateFile("controller", name, tmpl)
}

func generateUsecase(name string) {
	tmpl := `package usecase

type {{.Name}}UseCase interface {
	Get() string
	Post() string
	Put() string
	Delete() string
}

type {{.LowerName}}UseCase struct {
	// Add your dependencies here
}

func New{{.Name}}UseCase() {{.Name}}UseCase {
	return &{{.LowerName}}UseCase{}
}

func (u *{{.LowerName}}UseCase) Get() string {
	// Implement your logic here
	return "Get {{.Name}}"
}

func (u *{{.LowerName}}UseCase) Post() string {
	// Implement your logic here
	return "Post {{.Name}}"
}

func (u *{{.LowerName}}UseCase) Put() string {
	// Implement your logic here
	return "Put {{.Name}}"
}

func (u *{{.LowerName}}UseCase) Delete() string {
	// Implement your logic here
	return "Delete {{.Name}}"
}
`
	generateFile("usecase", name, tmpl)
}

func generateRouterWithoutCRUD(name string) {
	tmpl := `package router

import (
	"github.com/gofiber/fiber/v2"
	"newserver/internal/controller"
	"newserver/internal/usecase"
)

func New{{.Name}}Router(app *fiber.App, useCase usecase.{{.Name}}UseCase) {
	ctrl := controller.New{{.Name}}Controller(useCase)
	{{.LowerName}} := app.Group("/{{.LowerName}}")

	// Add your custom routes here
}
`
	generateFile("router", name, tmpl)
}

func generateControllerWithoutCRUD(name string) {
	tmpl := `package controller

import (
	"github.com/gofiber/fiber/v2"
	"newserver/internal/usecase"
)

type {{.Name}}Controller struct {
	useCase usecase.{{.Name}}UseCase
}

func New{{.Name}}Controller(useCase usecase.{{.Name}}UseCase) *{{.Name}}Controller {
	return &{{.Name}}Controller{useCase: useCase}
}

// Add your custom controller methods here
`
	generateFile("controller", name, tmpl)
}

func generateUsecaseWithoutCRUD(name string) {
	tmpl := `package usecase

type {{.Name}}UseCase interface {
	// Define your custom usecase methods here
}

type {{.LowerName}}UseCase struct {
	// Add your dependencies here
}

func New{{.Name}}UseCase() {{.Name}}UseCase {
	return &{{.LowerName}}UseCase{}
}

// Implement your custom usecase methods here
`
	generateFile("usecase", name, tmpl)
}

func generateRepository(name string) {
	tmpl := `package repository

import (
	"github.com/jmoiron/sqlx"
)

type {{.Name}}Repository interface {
	// Define your repository methods here
}

type {{.LowerName}}Repository struct {
	db *sqlx.DB
}

func New{{.Name}}Repository(db *sqlx.DB) {{.Name}}Repository {
	return &{{.LowerName}}Repository{db: db}
}
`
	generateFile("repository", name, tmpl)
}

func generateFile(dir, name, tmpl string) {
	lowerName := toLowerCamelCase(name)

	data := struct {
		Name      string
		LowerName string
	}{
		Name:      toUpperCamelCase(name),
		LowerName: lowerName,
	}

	t, err := template.New("file").Parse(tmpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	dirPath := fmt.Sprintf("internal/%s", dir)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	filePath := fmt.Sprintf("%s/%s.go", dirPath, lowerName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	err = t.Execute(file, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}

func toLowerCamelCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(append([]byte(strings.ToLower(string(s[0]))), s[1:]...))
}

func toUpperCamelCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(append([]byte(strings.ToUpper(string(s[0]))), s[1:]...))
}