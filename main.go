package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myblog/models"
	_ "myblog/routers"
)

func init() {
	models.CreateDB()
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.SessionOn = true
	beego.HttpPort = 9000
	beego.HttpAddr = "0.0.0.0"
	beego.Run()
}
