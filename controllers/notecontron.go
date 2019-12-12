package controllers

import (
	"bolg/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type NoteController struct {
	beego.Controller
}

type NewNoteController struct{
	beego.Controller
}
type DeleNoteController struct {
	beego.Controller
}

func (this *NoteController)Get() {
	this.TplName="editpage.html"
}

func (this *NoteController)Edit()  {
	//获取文章内容
	articlename:=this.GetString("articlename")
	content:=this.GetString("content")
	fmt.Println(articlename,content)
	if articlename==""{
		this.Data["notemsg"]="请输入文章标题"
		return
	}else if content== ""{
		this.Data["notemsg"]="请输入文章内容"
		return
	}
	//获取orm对象
	o:=orm.NewOrm()
	//插入数据库
	articcontent:=models.Note{}
	articcontent.Content=content
	articcontent.Summary=articlename
	//fmt.Println(articcontent)
	_,err:=o.Insert(&articcontent)
	if err!=nil{
		fmt.Println("插入数据错误")
		this.Data["notemsg"]="插入数据库出错"
		return
	}
	//this.TplName="home.html"
	this.Redirect("/home",302)
}

func (this *DeleNoteController)Delete()  {
	id,err:=this.GetInt("id")
	if err!=nil{
		fmt.Println("查询文章id出现错误",err)
		return
	}
	o:=orm.NewOrm()
	arti:=models.Note{Id:id}
	err=o.Read(&arti)
	if err!=nil{
		fmt.Println("查询有问题：",err)
		return
	}
	o.Delete(&arti)
	this.Redirect("/home",302)
}



