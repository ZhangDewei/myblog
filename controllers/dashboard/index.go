package dashboard

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"myblog/lib"
	"myblog/models"
)

type IndexDBController struct {
	beego.Controller
}

func ValidMaster(ctx *context.Context) ([]*models.User, bool) {
	user, err := lib.GetUserByCookie(ctx)
	if err == false {
		return user, false
	}
	if user[0].Super == 0 {
		return user, false
	}
	return user, true
}

func (this *IndexDBController) Get() {
	user, err := ValidMaster(this.Ctx)
	if err == false {
		this.Redirect("/", 301)
		return
	}
	this.Data["username"] = user[0].Name
	this.TplNames = "dashboard/index.html"
	this.Layout = "dashboard/layout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "dashboard/index_head.html"
	this.LayoutSections["Script"] = "dashboard/index_script.html"
}
