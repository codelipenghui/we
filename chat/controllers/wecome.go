package controllers

import "github.com/astaxie/beego"

type WelcomeController struct {
	beego.Controller
}

func (c *WelcomeController) Get() {
	c.TplName = "welcome.html"
	c.Render()
}
