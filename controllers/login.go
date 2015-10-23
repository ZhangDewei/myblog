package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myblog/contains"
	"myblog/lib"
	"myblog/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	_, err := lib.GetUserByCookie(this.Ctx)
	if err == true {
		this.Redirect("/", 301)
		return
	}
	this.TplNames = "login.html"
	this.Layout = "layout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "login_head.html"
	this.LayoutSections["Script"] = "login_script.html"
}

func (this *LoginController) CheckPwd() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	var result []*models.User
	o := orm.NewOrm()
	o.QueryTable("user").Filter("Name", username).Filter("Password", models.HashPassword(password)).One(&result)
	if len(result) == 1 {
		this.Data["json"] = true
	} else {
		this.Data["json"] = false
	}
	this.ServeJson()
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	remenber := this.Input().Get("remenber")
	if username == "" || password == "" {
		this.Redirect("/", 302)
		return
	}
	var user []*models.User
	o := orm.NewOrm()
	o.QueryTable("user").Filter("Name", username).Filter("Password", models.HashPassword(password)).One(&user)

	if len(user) == 0 {
		this.Redirect("/login", 302)
		return
	}
	fmt.Println(user)
	if remenber == "true" {
		lib.SetCookie(this.Ctx, username, password, contains.MiddleCookieTime)
	} else {
		lib.SetCookie(this.Ctx, username, password, contains.SmallCookieTime)
	}
	this.Redirect("/", 301)
}

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	_, err := lib.GetUserByCookie(this.Ctx)
	if err == true {
		lib.SetCookie(this.Ctx, "", "", -1)
	}
	this.Redirect("/", 301)
}
