package controllers

import "github.com/astaxie/beego"

type JoinController struct {
	beego.Controller
}

func (c *JoinController) Join() {
	username := c.GetString("username")
	c.Redirect("/"+username, 302)
	return
}
