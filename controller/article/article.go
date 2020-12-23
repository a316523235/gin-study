package article

import (
	"gin-study/lib"
	"gin-study/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func List(g *gin.Context) {
	ext := lib.BuildGinExt(g)
	limit, start := ext.GetLimitInfo()
	var res []models.Article
	err := lib.GetMysql().Where("is_del=0").Limit(limit, start).Find(&res)
	if err != nil {
		ext.Fault("get list error page: " + strconv.Itoa(limit) + " pagesize :" + strconv.Itoa(start))
		return
	}
	ext.Success(res, "ok")
}

func Info(g *gin.Context) {
	ext := lib.BuildGinExt(g)
	idStr := g.DefaultQuery("id", "")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ext.Fault("params error")
		return
	}
	var res models.Article
	isOk, err := lib.GetMysql().Id(id).Where("is_del=0").Get(&res)
	if err != nil || !isOk {
		ext.Fault404("404")
		return
	}
	ext.Success(res, "ok")
}
