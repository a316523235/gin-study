package lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

var mysqlEn *xorm.Engine

func GetMysql() *xorm.Engine {
	if mysqlEn == nil {
		dataSource := "root:root@tcp(127.0.0.1:3306)/gin_study?charset=utf8"
		en, err := xorm.NewEngine("mysql", dataSource)
		if err != nil {
			_, _ = fmt.Fprintln(gin.DefaultWriter, "mysql error : "+err.Error())
			panic(err.Error())
		}
		mysqlEn = en
	}

	return mysqlEn
}
