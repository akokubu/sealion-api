package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// routing
	e.GET("/tasks", indexTask)
	e.GET("/tasks/:id", showTask)

	e.POST("/tasks", createTask)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}

func indexTask(c echo.Context) error {
	var tasks []Task
	db.Find(&tasks)

	jsonObject := map[string][]Task{"tasks": tasks}
	return c.JSON(http.StatusOK, jsonObject)
}

func showTask(c echo.Context) error {
	var task Task
	db.First(&task, c.Param("id"))
	return c.JSON(http.StatusOK, task)
}

func createTask(c echo.Context) error {
	task := Task{Title: "ABCDE", Project: "P Q R", Done: false}
	db.Create(&task)
	return c.JSON(http.StatusOK, task)
}
