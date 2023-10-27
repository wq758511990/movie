package routers

import (
	"github.com/gin-gonic/gin"
	"todo-pro/handlers/movieHandlers"
)

func InitMovieRouter(g *gin.Engine) {
	movieRouter := g.Group("/api/v1/movie")
	{
		movieRouter.GET("/", movieHandlers.MovieGet)
	}
}
