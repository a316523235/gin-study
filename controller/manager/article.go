package manager

import (
	"gin-study/lib"
	"gin-study/models"
	"github.com/gin-gonic/gin"
)

func ArticleAdd(g *gin.Context)  {
	ext := lib.BuildGinExt(g)
	var article models.Article
	err := g.ShouldBind(&article)
	if err != nil {
		ext.Fault("error params")
		return
	}
	_, err = lib.GetMysql().Insert(&article)
	if err != nil {
		ext.Fault("save error")
		return
	}
	ext.Success(article, "ok")
}

func ArticleEdit(g *gin.Context)  {
	ext := lib.BuildGinExt(g)
	var article models.Article
	err := g.ShouldBind(&article)
	if err != nil || article.Id == 0 {
		ext.Fault("error params")
		return
	}
	_, err = lib.GetMysql().Id(article.Id).Update(&article)
	if err != nil {
		ext.Fault("save error")
		return
	}
	ext.Success(article, "ok")
}