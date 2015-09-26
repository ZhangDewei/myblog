package lib

import (
	//"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"myblog/models"
)

func CheckAccount(ctx *context.Context) bool {
	_, unameErr := ctx.Request.Cookie("uname")
	_, pwdErr := ctx.Request.Cookie("pwd")
	if unameErr != nil || pwdErr != nil {
		return false
	} else {
		return true
	}
}

func SetCookie(ctx *context.Context, username, password string, expire int) {
	ctx.SetCookie("uname", username, expire)
	ctx.SetCookie("pwd", password, expire)
}

func GetUserByCookie(ctx *context.Context) (interface{}, bool) {
	uk, unameErr := ctx.Request.Cookie("uname")
	pk, pwdErr := ctx.Request.Cookie("pwd")
	if unameErr != nil || pwdErr != nil {
		return "", false
	}
	o := orm.NewOrm()
	var result []*models.User
	o.QueryTable("user").Filter("Name", uk).Filter("Password", pk).One(&result)
	if len(result) > 0 {
		return result, true
	} else {
		return "", false
	}
}
