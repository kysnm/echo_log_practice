package main

import (
	"log"
	"net/http"
	"os"

	"github.com/comail/colog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"time"
)

func GetTest() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Echo().Logger.Debug("debug at handler")
		return c.JSON(http.StatusOK, "hello world!")
	}
}

func main() {
	e := echo.New()

	file, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	log.SetOutput(file)

	e.Debug = true
	e.Logger.SetOutput(file)

	log.Println("foo")

	log.Println("debug: from colog")

	var out = colog.NewCoLog(file, "", 0)
	out.SetFormatter(&colog.JSONFormatter{
		TimeFormat: time.RFC3339,
		Flag:       log.Lshortfile,
	})

	log.SetOutput(out)

	log.Println("debug: colog in json format")

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: file,
	}))

	e.GET("/test", GetTest())

	e.Start(":8000")
}
