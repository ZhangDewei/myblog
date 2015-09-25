package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

type HomeController struct{
	beego.Controller
}

func (this *HomeController) Get(){
	this.Data["Name"] = "dewei"
    name := this.GetSession("name")
    this.Data["session"] = name
	this.TplNames = "index.html"
    this.Layout = "layout.html"
}

type VideoController struct{
	beego.Controller
}

func (this *VideoController) Get(){
	this.TplNames = "myvideo.html"
}