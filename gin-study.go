package main

import (
	"gin-study/controller"
	"gin-study/controller/article"
	"gin-study/controller/manager"
	"gin-study/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	bind(r)
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}

func bind(r *gin.Engine) {

	r.LoadHTMLGlob("views/*")
	r.GET("/", controller.Index)
	r.GET("/ping", controller.Ping)
	r.GET("/redirect", controller.Redirect)
	r.GET("/redirect2", controller.Redirect2)
	r.GET("/jsonp", controller.Jsonp)

	r.GET("/article/list", article.List)
	r.GET("/article/info", article.Info)

	r.POST("/manager/login", manager.Login)
	group := r.Group("/manager")
	group.Use(middleware.AdminAuth())
	group.GET("/auth/check", manager.AuthCheck)
	group.POST("/article/add", manager.ArticleAdd)
	group.PUT("/article/edit", manager.ArticleEdit)
}
