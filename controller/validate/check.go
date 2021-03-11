package validate

import (
	"fmt"
	"gin-study/lib"
	"github.com/gin-gonic/gin"
	"time"
)

type Check1 struct {
	Age int `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func V1(g * gin.Context)  {
	var ck Check1
	ext := lib.BuildGinExt(g)
	if err := g.ShouldBind(&ck); err != nil {
		ext.Success(fmt.Sprint(err), "xx")
		return
	}
	ext.Success(ck, "ok")
}
