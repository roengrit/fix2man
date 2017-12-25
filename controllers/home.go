package controllers

import (
	"github.com/astaxie/beego"
)

//AppController _
type AppController struct {
	beego.Controller
}

//Get _
func (c *AppController) Get() {

	secret := beego.AppConfig.String("secret")
	val, ok := c.Ctx.GetSecureCookie(secret, "fixman")
	_ = val
	if !ok {
		c.Ctx.Redirect(302, "/auth")
	}
	c.Layout = "layout.html"
	c.TplName = "main/index.html"
	c.Render()
}
