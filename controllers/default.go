package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblog/lib"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["Name"] = "dewei"
	//name := this.GetSession("name")
	//this.Data["session"] = name
	this.TplNames = "index.html"
	this.Layout = "layout.html"
	//this.LayoutSessions = make(map[string]string)
	//this.LayoutSections["HtmlHead"] = "home_head.html"
	//this.LayoutSections["Script"] = "home_script.html"
	user, err := lib.GetUserByCookie(this.Ctx)
	if err == true {
		fmt.Println(user)
	}

}

type VideoController struct {
	beego.Controller
}

func (this *VideoController) Get() {
	this.TplNames = "myvideo.html"
}
