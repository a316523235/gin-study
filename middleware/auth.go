package middleware

import (
	"encoding/json"
	"fmt"
	"gin-study/conf"
	"gin-study/lib"
	"gin-study/models"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func AdminAuth() gin.HandlerFunc {
	return func(g *gin.Context) {
		ext := lib.BuildGinExt(g)
		adminToken, err := g.Cookie("admin_token")
		if err != nil || adminToken == "" {
			ext.Fault("un login")
			g.Abort()
			return
		}
		key := fmt.Sprintf(conf.AdminTokenKey, adminToken)
		reply, err := lib.GetRedis().Do("get", key)
		adminInfoStr, err := redis.String(reply, err)
		var adminInfo models.Admin
		err = json.Unmarshal([]byte(adminInfoStr), &adminInfo)
		if err != nil || adminInfo.Id == 0 {
			ext.Fault("expire")
			g.Abort()
			return
		}
		g.Set("adminInfo", adminInfo)
	}
}
