package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"os"
    "os/exec"
    "path/filepath"
	"strings"
	"app/models"
)

func getPath() string{
	file, _ := exec.LookPath(os.Args[0])
    path, _ := filepath.Abs(file)
	result := strings.Trim(path, "myblog")
	return result
}

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

func (this *RegisterController) Post(){
	_,h,img_err:=this.GetFile("user-header")
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	email := this.Input().Get("email")
	gender := this.Input().Get("gender")
	age := this.Input().Get("age")
	var path = getPath() + "static/img/head/"
	var img_header string = ""
	if img_err == nil{
		img_header = fmt.Sprint(path, h.Filename)
		this.SaveToFile("user-header", img_header)
	}
	
	o := orm.NewOrm()
	o.Using("default")
	
	var userObj *models.User = new(models.User)
	userObj.Name = username
	userObj.Password = password
	userObj.Email = email
	userObj.Gender = gender
	userObj.Age = age
	userObj.Avator = img_header
	
	o.Insert(userObj)
	
	
	this.Redirect("/", 302)
}