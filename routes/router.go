package routes

import (
	"crypto/subtle"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/SakaiTaka23/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func ParseTemplates() *template.Template {
	templ := template.New("")
	err := filepath.Walk("./views", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return templ
}

func Routes() {
	_ = godotenv.Load(".././.env")
	// ADMIN_NAME := os.Getenv("ADMIN_NAME")
	// ADMIN_PASS := os.Getenv("ADMIN_PASS")
	// fmt.Print("admin_name")
	// fmt.Print(ADMIN_NAME, ADMIN_PASS)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Static("/", "assets")

	renderer := &TemplateRenderer{
		templates: ParseTemplates(),
	}
	e.Renderer = renderer

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("admin")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("pass")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/post", handler.PostForm)

	g.POST("/post", func(c echo.Context) error {
		return c.String(http.StatusOK, "/post POST")
	})

	e.GET("/article", func(c echo.Context) error {
		return c.String(http.StatusOK, "/article")
	})

	e.GET("/article/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/article/:id")
	})

	e.GET("/tag", func(c echo.Context) error {
		return c.String(http.StatusOK, "/tag")
	})

	e.GET("/tag/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/tag:id")
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//  ルート書き出し
	// data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	// _ = ioutil.WriteFile("routes.json", data, 0644)

	e.Logger.Fatal(e.Start(":8082"))
}
