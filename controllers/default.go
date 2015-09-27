package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"myblog/lib"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	//name := this.GetSession("name")
	//this.Data["session"] = name
	this.TplNames = "index.html"
	this.Layout = "layout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "index_head.html"
	this.LayoutSections["Script"] = "index_script.html"
	user, err := lib.GetUserByCookie(this.Ctx)
	if err == true {
		this.Data["uname"] = user[0].Name
	}
}

type VideoController struct {
	beego.Controller
}

func (this *VideoController) Get() {
	this.TplNames = "myvideo.html"
}
