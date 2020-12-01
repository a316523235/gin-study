package article

import (
	"gin-study/lib"
	"github.com/gin-gonic/gin"
)

func list(g gin.Context)  {
	ext := lib.BuildGinExt(g)
	page, pagesize := ext.GetPageInfo()
}
