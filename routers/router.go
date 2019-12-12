package routers

import (
	"bolg/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/cookie",&controllers.RegisterController{})
    beego.Router("/", &controllers.ShowloginController{},"GET:Get;POST:Showlogin")
    beego.Router("/home",&controllers.MainController{},"*:Showhome")
    beego.Router("/register",&controllers.RegisterController{},"GET:Showregister;POST:Handleregister")
    beego.Router("/editpage",&controllers.NoteController{},"GET:Get;POST:Edit")
    beego.Router("/blog",&controllers.ShowblogController{},"get:Showblog")
    beego.Router("/delete",&controllers.DeleNoteController{},"get:Delete")
	}
