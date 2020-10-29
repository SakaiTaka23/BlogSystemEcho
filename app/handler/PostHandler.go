package handler

import (
	"github.com/labstack/echo"
)

func PostForm(c echo.Context) error {
	return c.File("admin/post-form.html")
}
