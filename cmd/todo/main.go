package main

import (
	"todo-cli/internal/handlers/todohdl"
	"todo-cli/internal/pkg/uidgen"
	"todo-cli/internal/repositories/todorepo"
	"todo-cli/internal/services/todosrv"

	"github.com/gin-gonic/gin"
)

func main() {
	todoRepository := todorepo.NewJSONRepository()
	todoService := todosrv.New(todoRepository, uidgen.New())
	todoHandler := todohdl.NewHTTPHandler(todoService)

	router := gin.New()

	router.GET("/todos", todoHandler.GetAll)
	router.GET("/todos/:id", todoHandler.GetById)
	router.POST("/todos", todoHandler.Create)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
