package main


import (
	"github.com/labstack/echo"
	"net/http"
)


func GetTest() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world!")
	}
}

func main() {
	e := echo.New()

	e.GET("/test", GetTest())

	e.Start(":8000")
}
