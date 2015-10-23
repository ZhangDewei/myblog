package dashboard

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myblog/models"
)

type VideoDBController struct {
	beego.Controller
}

func (this *VideoDBController) Get() {
	user, err := ValidMaster(this.Ctx)
	if err == false {
		this.Redirect("/", 301)
		return
	}
	var videoObjs []*models.Video
	var video *models.Video = new(models.Video)
	o := orm.NewOrm()
	o.QueryTable(video).RelatedSel().All(&videoObjs)

	this.Data["username"] = user[0].Name
	this.Data["user_id"] = user[0].Id
	this.Data["videos"] = videoObjs
	this.TplNames = "dashboard/video.html"
	this.Layout = "dashboard/layout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "dashboard/video_head.html"
	this.LayoutSections["Script"] = "dashboard/video_script.html"
}

func (this *VideoDBController) Post() {
	userId := this.Input().Get("user_id")
	url := this.Input().Get("url")
	comFrom := this.Input().Get("comrrom")
	des := this.Input().Get("desc")

	user := new(models.User)
	var userObj []*models.User
	o := orm.NewOrm()
	o.QueryTable(user).Filter("Id", userId).One(&userObj)
	var videoObj *models.Video = new(models.Video)
	videoObj.Add(userObj[0], url, comFrom, des)
	this.Redirect("/dashboard/video", 301)
}
