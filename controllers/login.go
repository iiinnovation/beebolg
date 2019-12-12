package controllers

import (
	"bolg/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ManiController struct {
	beego.Controller
}
type ShowloginController struct {
	beego.Controller
}
type ShowblogController struct {
	beego.Controller
}

func (c *ShowloginController)Get()  {
	c.TplName="index.html"

}
//博客主界面
func (c *MainController)Showhome() {
	username:=c.GetSession("username")
	if username==nil{
		c.Redirect("/",302)
		return
	}
	//获取orm对象
	o:=orm.NewOrm()
	//新建一个数组存放articles
	var articles []models.Note
	//读取全部articles的数据
	_,err:=o.QueryTable("Note").All(&articles)
	if err!=nil{
		fmt.Println("查询内容错误")
		return
	}
	fmt.Println(articles)
	//写到网页
	c.Data["articles"]=articles
	//返回一个带数据到网页
	c.TplName="home.html"

}

func (this *MainController)Gethomepage()  {
	this.TplName="home.html"
}

//注册页面

func (this *ShowloginController)Showlogin()  {
	//获取浏览器数据
	uname:=this.GetString("username")
	pwd:=this.GetString("password")
	if uname=="" {
		this.Data["errmsg1"] = "用户名不能为空"
	}else if pwd==""{
		this.Data["errmsg"]="密码不能为空"
		fmt.Printf("用户和密码不能为空")
		this.TplName="index.html"
		return
	}
	//查询用户名和密码
	//获取orm对象
	o:=orm.NewOrm()
	//获取models对象
	var user models.User
	//查询用户名
	user.UserName=uname
	err:=o.Read(&user,"Username")
	if err!=nil{
		fmt.Printf("用户名错误")
		this.TplName="index.html"
		this.Data["errmsg"]="用户名错误"
		return
	}
	//查询密码
	if user.Password!=pwd{
		beego.Info("密码错误")
		this.TplName="index.html"
		this.Data["errmsg"]="密码错误"
		return
	}
	//返回主页面
	this.SetSession("username",user)
	this.Redirect("/home",302)

}


//博客详情页面
func (this *ShowblogController)Showblog()  {
	//获取id
	id,err:=this.GetInt("id")
	fmt.Println(id)
		if err!=nil{
			fmt.Println("查询id出现错误",err)
			return
		}
	//根据id查询内容
	o:=orm.NewOrm()
	arti:=models.Note{Id:id}
	err=o.Read(&arti)
	if err!=nil{
		fmt.Println("读取ID出现错误",err)
	}
	//返回数据
	this.Data["article"]=arti
	//打印带数据的页面
	this.TplName="blog.html"
}