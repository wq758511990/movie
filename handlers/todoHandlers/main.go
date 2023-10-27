package todoHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-pro/common"
	"todo-pro/globals"
	"todo-pro/models"
)

func CreateTodo(ctx *gin.Context) {
	completed, _ := strconv.Atoi(ctx.PostForm("completed"))
	todo := models.TodoModel{Title: ctx.PostForm("title"), Completed: completed}
	globals.DB.Save(&todo)
	data := map[string]uint{
		"resourceId": todo.ID,
	}
	ctx.JSON(http.StatusCreated, common.OK.WithMessage("Todo item created successfully!").WithData(data))
}

func FetchAllTodo(ctx *gin.Context) {
	var todos []models.TodoModel
	var _todos []models.TransformedTodo
	globals.DB.Find(&todos)
	if len(todos) <= 0 {
		ctx.JSON(http.StatusNotFound, common.Err.WithMessage("no todo lists found!"))
		return
	}
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, models.TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}
	ctx.JSON(http.StatusOK, common.OK.WithData(_todos))
}

func FetchSingleTodo(ctx *gin.Context) {
	var todo models.TodoModel
	todoID := ctx.Param("id")
	globals.DB.First(&todo, todoID)
	if todo.ID == 0 {
		ctx.JSON(http.StatusNotFound, common.Err.WithMessage("no todo found"))
		return
	}
	completed := false
	if todo.Completed == 1 {
		completed = true
	}
	_todo := models.TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
	ctx.JSON(http.StatusOK, common.OK.WithData(_todo))
}

func UpdateTodo(ctx *gin.Context) {
	var todo models.TodoModel
	todoID := ctx.Param("id")
	globals.DB.First(&todo, todoID)
	if todo.ID == 0 {
		ctx.JSON(http.StatusNotFound, common.Err.WithMessage("no todo found!"))
		return
	}
	globals.DB.Model(&todo).Update("title", ctx.PostForm("title"))
	completed, _ := strconv.Atoi(ctx.PostForm("completed"))
	globals.DB.Model(&todo).Update("completed", completed)
	ctx.JSON(http.StatusOK, common.OK.WithMessage("Todo Update successfully!"))
}

func DeleteTodo(ctx *gin.Context) {
	todo := models.TodoModel{}
	globals.DB.First(&todo, ctx.Param("id"))

	if todo.ID == 0 {
		ctx.JSON(http.StatusNotFound, common.Err.WithMessage("no todo found"))
		return
	}

	globals.DB.Delete(&todo)
	ctx.JSON(http.StatusOK, common.OK.WithMessage("todo deleted successfully!"))
}
