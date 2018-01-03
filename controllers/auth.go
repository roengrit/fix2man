package controllers

import (
	h "fix2man/helps"
	"fix2man/models"
	"html/template"
	"net/http"

	"github.com/astaxie/beego"
)

//AuthController _
type AuthController struct {
	beego.Controller
}

//LogoutController _
type LogoutController struct {
	beego.Controller
}

//Get to view login
func (c *AuthController) Get() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["username"] = "logon.firstclass@gmail.com"
	c.Data["title"] = "เข้าสู่ระบบเพื่อเริ่มการทำงาน"
	c.TplName = "auth/login.html"
	c.Render()
}

//Post to login
func (c *AuthController) Post() {
	usernameForm := c.GetString("username")
	passwordForm := c.GetString("password")

	if ok, err := models.Login(usernameForm, passwordForm); ok {
		user, _ := models.GetUserByUserName(usernameForm)
		if ok, err = h.KeepLogin(c.Ctx.ResponseWriter, usernameForm, user.Roles.ID, user.Branch.ID); ok == true {
			c.Ctx.Redirect(http.StatusFound, "/")
		} else {
			c.Data["error"] = err
		}
	} else {
		c.Data["error"] = err
	}

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["username"] = c.GetString("username")
	c.Data["title"] = "เข้าสู่ระบบเพื่อเริ่มการทำงาน"
	c.TplName = "auth/login.html"
	c.Render()
}

//Get to logout
func (c *LogoutController) Get() {
	h.LogOut(c.Ctx.ResponseWriter)
	c.Ctx.Redirect(http.StatusFound, "/auth")
}
