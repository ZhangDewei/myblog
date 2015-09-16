package main

import (
	_ "myblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myblog/models"
)

func init(){
	models.CreateDB()
	orm.RunSyncdb("default", false, true) 
}

func main() {
	beego.Run()
}

