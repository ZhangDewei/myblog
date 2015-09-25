package controllers

import (
	"github.com/astaxie/beego"
)


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