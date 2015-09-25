package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"os"
    "os/exec"
    "path/filepath"
	"strings"
	"myblog/models"
	"strconv"
)

func getPath() string{
	file, _ := exec.LookPath(os.Args[0])
    path, _ := filepath.Abs(file)
	result := strings.Trim(path, "myblog")
	return result
}

func existUsername(name string) bool{
	o := orm.NewOrm()
    var result [] *models.User
    o.QueryTable("user").Filter("Name", name).One(&result)
    // result 我定义了一个数组 里边每个user对象都可以直接调用 如 .Id, .Name
    if len(result) > 0{
        return false
    }else{
        return true
    }
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

func (this *RegisterController) CheckUserExist(){
    username := this.Input().Get("username")
    result := make(map[string] bool)
    valid_user := existUsername(username)  
    result["content"] = valid_user
    this.Data["json"] = result
    fmt.Println(valid_user)
    this.ServeJson()
}

func (this *RegisterController) Post(){
    username := this.Input().Get("username")    
	_,h,img_err:=this.GetFile("user-header")
	
	password := this.Input().Get("password")
	email := this.Input().Get("email")
	gender := this.Input().Get("gender")
	age := this.Input().Get("age")
    
    valid_user := existUsername(username)    
    if valid_user == false{
        this.Redirect("/register", 302)
        return
    }
	
	var path = getPath() + "static/img/head/"
	var img_header string = ""
	if img_err == nil{
		img_header = fmt.Sprint(path, h.Filename)
		this.SaveToFile("user-header", img_header)
	}
	
	int_age, _ := strconv.Atoi(age)
	int_gender, _ := strconv.Atoi(gender)
	
	o := orm.NewOrm()
	o.Using("default")
	
	var userObj *models.User = new(models.User)
	userObj.Name = username
	userObj.Password = password
	userObj.Email = email
	userObj.Gender = int32(int_gender)
	userObj.Age = int32(int_age)
	userObj.Avator = img_header
	
	o.Insert(userObj)
	
	this.Redirect("/", 302)
}