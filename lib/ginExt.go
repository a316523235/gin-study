package lib

import (
	"gin-study/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ginExt struct {
	g         *gin.Context
	adminInfo *models.Admin
}

func BuildGinExt(g *gin.Context) *ginExt {
	return &ginExt{g: g}
}

func (x *ginExt) Success(data interface{}, message string) {
	x.g.JSON(200, gin.H{
		"data":    data,
		"message": message,
	})
}

func (x *ginExt) Fault(message string) {
	x.g.JSON(400, gin.H{
		"data":    nil,
		"message": message,
	})
}

func (x *ginExt) Fault404(message string) {
	x.g.JSON(404, gin.H{
		"data":    nil,
		"message": message,
	})
}

func (x *ginExt) GetPageInfo() (page, pageSize int) {
	pageStr, pageSizeStr := x.g.DefaultQuery("page", "1"), x.g.DefaultQuery("pagesize", "20")
	page, _ = strconv.Atoi(pageStr)
	pageSize, _ = strconv.Atoi(pageSizeStr)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 && pageSize > 50 {
		pageSize = 20
	}
	return page, pageSize
}

func (x *ginExt) GetLimitInfo() (limit, start int) {
	page, pagesize := x.GetPageInfo()
	limit, start = page*pagesize, (page-1)*pagesize
	return
}

func (x *ginExt) GetAdminInfo() models.Admin {
	if x.adminInfo == nil {
		str, _ := x.g.Get("adminInfo")
		xx := str.(models.Admin)
		x.adminInfo = &xx
	}
	return *x.adminInfo
}
