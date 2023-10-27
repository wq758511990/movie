package initialize

import (
	"github.com/gin-gonic/gin"
	"todo-pro/routers"
)

func InitRouters(r *gin.Engine) {
	routers.InitMovieRouter(r)
}
