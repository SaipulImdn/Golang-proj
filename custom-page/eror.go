package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError)
		}

		errPage := fmt.Sprintf("%d.html", report.Code)
		if err := c.File(errPage); err != nil {
			c.HTML(report.Code, "ERROR")
		}
	}

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, true)
	})

	e.Logger.Fatal(e.Start(":9000"))
}
