package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
	"myblog/lib"
)

type User struct{
	Id int64
	Name string `orm:"default(null);size(50);index"`
	Email string `orm:"default(null);size(100)"`
	Password string `orm:"default(null);size(200)"`
	Avator string `orm:"default(null);size(250)"`
	Gender int32 `orm:"default(0)"`
	Age int32 `orm:"default(0)"`
	Status string `orm:"default(D);size(2)"`
}

func (user *User) toString() string{
	return fmt.Sprint("user %s", user.Name)
}

func (user *User) Add(username string, email string, password string, avator string, gender int, age int){
	hashPassword := lib.HashPassword(password)
	
	user.Name = username
	user.Email = email
	user.Password = hashPassword
	user.Avator = avator
	user.Gender = int32(gender)
	user.Age = int32(age)
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(user)
}

type Article struct{
	Id int64
	User *User `orm:"rel(fk)"`
	Title string `orm:"default(null);size(100)"`
	Content string `orm:"default(null);size(1000)"`
	Created time.Time
}

type Replay struct{
	Id int64
	User *User `orm:"rel(fk)"`
	Article *Article `orm:"rel(fk);index"`
	Content string `orm:"default(null);size(1000)"`
	Created time.Time
}

func CreateDB(){
	orm.RegisterModel(new(User), new(Article), new(Replay))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/myblog?charset=utf8")
}