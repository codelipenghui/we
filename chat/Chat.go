package main

import (
	"github.com/astaxie/beego"
	"we/chat/controllers"
)

func main() {
	beego.Router("/", &controllers.WelcomeController{})
	beego.Router("/join", &controllers.JoinController{}, "post:Join")
	beego.Run()
}