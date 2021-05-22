package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-embed-sample/filebox"
	"net/http"
)

func EchoStart(port int32) {
	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		const indexHtml = "index.html"
		if filebox.Exists(indexHtml) {
			blob, contentType := filebox.GetFile(indexHtml)
			return c.Blob(http.StatusOK, contentType, blob)
		}

		return c.String(http.StatusNotFound, "Not Found")
	})

	e.GET("/:filename", func(c echo.Context) error {
		fn := c.Param("filename")
		fmt.Println("Embed: ", fn)

		if filebox.Exists(fn) {
			blob, contentType := filebox.GetFile(fn)
			return c.Blob(http.StatusOK, contentType, blob)
		}

		return c.String(http.StatusNotFound, "Not Found")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
