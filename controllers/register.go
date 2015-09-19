package controllers

import (
	"github.com/astaxie/beego"
)

type RegisterController struct{
	beego.Controller
}

func (this *RegisterController) Get(){
	this.TplNames = "register.html"
	this.Layout = "layout.html"
	this.LayoutSections = make(map[string] string)
	this.LayoutSections["HtmlHead"] = "register_head.html"
	this.LayoutSections["Script"] = "register_script.html"
}