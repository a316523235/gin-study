package manager

import (
	"encoding/json"
	"fmt"
	"gin-study/conf"
	"gin-study/lib"
	"gin-study/models"
	"github.com/gin-gonic/gin"
)

func Login(g *gin.Context) {
	ext := lib.BuildGinExt(g)
	username, password := g.PostForm("username"), g.PostForm("pwd")
	if username == "" || password == "" {
		ext.Fault("params empty")
		return
	}
	encodePwd := lib.EncodePwd(password)
	var admin models.Admin
	isOk, err := lib.GetMysql().Where("username=? and pwd=?", username, encodePwd).Get(&admin)
	if err != nil || !isOk {
		ext.Fault("not user")
		return
	}
	admin.Pwd = ""
	bt, _ := json.Marshal(admin)
	adminInfo := string(bt)
	token := lib.EncodeToken(username, encodePwd)
	key := fmt.Sprintf(conf.AdminTokenKey, token)
	_, _ = lib.GetRedis().Do("set", key, adminInfo)
	_, _ = lib.GetRedis().Do("expire", key, 60*60*24)
	g.SetCookie("admin_token", token, 60*60*24, "", "", false, true)
	ext.Success(admin, "ok")
}

func AuthCheck(g *gin.Context) {
	ext := lib.BuildGinExt(g)
	admin := ext.GetAdminInfo()
	ext.Success(admin, "ok")
}
