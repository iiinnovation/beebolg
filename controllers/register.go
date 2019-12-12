package controllers

import (
	"bolg/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}
type XsrfController struct {
	beego.Controller
}
func (this *RegisterController)Showregister() {
	this.TplName="register.html"
}
//获取注册页面数据
func (this *RegisterController) Handleregister()  {
	//获取用户名
	Name:=this.GetString("username")
	//获取用户密码
	pwd:=this.GetString("password")
	fmt.Println(Name,pwd)
	if Name==""||pwd==""{
		beego.Info("用户和密码不能为空")
		this.TplName="register.html"
		this.Data["errmsg"]="注册用户和密码不能为空"
		return
	}
	//插入数据库
	//获取orm对象
	o:=orm.NewOrm()
	//获取插入对象
	user:=models.User{}
	//插入数值
	user.UserName=Name
	user.Password=pwd
	_,err:=o.Insert(&user)
	if err!=nil {
		this.TplName = "register.html"
		this.Data["errmsg"] = "插入数据错误"
		return
	}

	//	重定向回登陆首页
	//this.Ctx.WriteString("测试成功！！！")
	this.TplName="index.html"
}
