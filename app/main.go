package main

import (
	"crypto/subtle"
	"html/template"
	"io"
	"net/http"
	"os"

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

func main() {
	_ = godotenv.Load()
	ADMIN_NAME := os.Getenv("ADMIN_NAME")
	ADMIN_PASS := os.Getenv("ADMIN_PASS")

	e := echo.New()

	e.Static("/", "assets")

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte(ADMIN_NAME)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(ADMIN_PASS)) == 1 {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/post", func(c echo.Context) error {
		return c.String(http.StatusOK, "/post GET")
	})

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
