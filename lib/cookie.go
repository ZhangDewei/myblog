package lib

import (
	"fmt"
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
	ctx.SetCookie("uname", username, expire, "/")
	ctx.SetCookie("pwd", password, expire, "/")
}

func GetUserByCookie(ctx *context.Context) ([]*models.User, bool) {
	uk, unameErr := ctx.Request.Cookie("uname")
	pk, pwdErr := ctx.Request.Cookie("pwd")
	var result []*models.User
	if unameErr != nil || pwdErr != nil {
		return result, false
	}
	o := orm.NewOrm()
	o.QueryTable("user").Filter("Name", uk.Value).Filter("Password", models.HashPassword(pk.Value)).One(&result)
	fmt.Println(result)
	if len(result) > 0 {
		return result, true
	} else {
		return result, false
	}
}
