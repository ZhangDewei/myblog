package dashboard

import (
	"github.com/astaxie/beego"
)

type IndexDBController struct {
	beego.Controller
}

func (this *IndexDBController) Get() {
	this.TplNames = "dashboard/index.html"
}
