package lib

import (
	"fmt"
	"gin-study/conf"
)

func EncodePwd(pwd string) string {
	return Md5(conf.PasswordHalt + pwd)
}

func EncodeToken(username, encodePwd string) string {
	str := fmt.Sprintf("%s.%s.%s", conf.TokenHalt, username, encodePwd)
	return Md5(str)
}
