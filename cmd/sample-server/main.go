package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nvvu/goast-demo/handlers"
)

func main() {

	r := gin.Default()

	r.POST("/", gin.WrapF(handlers.CreateUser_hdl))
	r.POST("/get-user", gin.WrapF(handlers.GetUser_hdl))

	r.Run()
}
