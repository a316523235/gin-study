package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
)

func Index(c *gin.Context) {
	data := render.Data{}
	c.HTML(200, "index.html", gin.H{
		"title": "test bind value",
		"data": data,
	})
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
}

func Redirect2(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}

func Jsonp(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{
		"foo": "bar",
	})
}
