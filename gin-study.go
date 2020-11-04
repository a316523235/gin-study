package main

import (
	"gin-study/controller"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	bind(r)
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}

func bind(r *gin.Engine)  {

	r.LoadHTMLGlob("views/*")
	r.GET("/", controller.Index)
	r.GET("/ping", controller.Ping)
	r.GET("/redirect", controller.Redirect)
	r.GET("/redirect2", controller.Redirect2)
	r.GET("/jsonp", controller.Jsonp)
}