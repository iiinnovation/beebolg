package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
type User struct {
	Id int
	Password string	`form:"Password"`
	UserName string	`form:"UserName"`

}
type Note struct {
	Id int
	UserName string
	Key string
	Content string
	Summary string
}
func init()  {
	//注册数据设置
	orm.RegisterDataBase("default","mysql","root:qwer110asd@tcp(127.0.0.1:3306)/bolg?charset=utf8")
	//映射models
	orm.RegisterModel(new(User),new(Note))
	//生成表
	orm.RunSyncdb("default",false,true)
}