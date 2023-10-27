package routers

import (
	"github.com/gin-gonic/gin"
	"todo-pro/handlers/todoHandlers"
)

func InitTodoRouter(g *gin.Engine) {
	todoRouters := g.Group("/api/v1/todos")
	todoRouters.POST("/", todoHandlers.CreateTodo)
	todoRouters.GET("/", todoHandlers.FetchAllTodo)
	todoRouters.GET("/:id", todoHandlers.FetchSingleTodo)
	todoRouters.PUT("/:id", todoHandlers.UpdateTodo)
	todoRouters.DELETE("/:id", todoHandlers.DeleteTodo)
}
