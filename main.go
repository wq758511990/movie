package main

import (
	"github.com/gin-gonic/gin"
	"todo-pro/initialize"
)

func main() {
	r := gin.Default()
	initialize.InitRouters(r)
	initialize.InitDB()

	r.Run(":8991")
}
