package todohdl

import (
	"todo-cli/internal/ports"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	todoService ports.TodoService
}

func NewHTTPHandler(todoService ports.TodoService) *HTTPHandler {
	return &HTTPHandler{
		todoService: todoService,
	}
}

func (h *HTTPHandler) GetAll(c *gin.Context) {
	todos, err := h.todoService.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, todos)
}

func (h *HTTPHandler) GetById(c *gin.Context) {
	todo, err := h.todoService.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, todo)
}

func (h *HTTPHandler) Create(c *gin.Context) {
	var req struct {
		Description string `json:"description"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.todoService.Create(req.Description)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"id": id})
}
