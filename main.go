package main

import (
	_ "bolg/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/mysql"
)


func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

