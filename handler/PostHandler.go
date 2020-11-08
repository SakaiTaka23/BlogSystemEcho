package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func PostForm(c echo.Context) error {
	data := struct {
		Title string
	}{
		Title: "PostForm",
	}
	return c.Render(http.StatusOK, "PostForm", data)
}
