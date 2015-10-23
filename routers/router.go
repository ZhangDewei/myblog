package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
	"myblog/controllers/dashboard"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/register/checkuser", &controllers.RegisterController{}, "post:CheckUserExist")
	beego.Router("/myvideo", &controllers.VideoController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/login/checkuser", &controllers.LoginController{}, "post:CheckPwd")
	beego.Router("/logout", &controllers.LogoutController{})

	beego.Router("/dashboard", &dashboard.IndexDBController{})
	beego.Router("/dashboard/video", &dashboard.VideoDBController{})
}
