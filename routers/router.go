package routers

import (
	"myblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
    beego.Router("/register/checkuser", &controllers.RegisterController{}, "post:CheckUserExist")
	beego.Router("/myvideo", &controllers.VideoController{})
	
}
