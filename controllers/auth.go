package controllers

import (
	h "fix2man/helps"
	"fix2man/models"
	"html/template"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
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
	c.Data["username"] = ""
	c.Data["title"] = "เข้าสู่ระบบเพื่อเริ่มการทำงาน"
	c.TplName = "auth/index.html"
	c.Render()
}

//Post to login
func (c *AuthController) Post() {

	c.Data["username"] = c.GetString("username")
	c.Data["title"] = "เข้าสู่ระบบเพื่อเริ่มการทำงาน"

	usernameForm := c.GetString("username")
	passwordForm := c.GetString("password")

	o := orm.NewOrm()
	user := models.Users{Username: usernameForm}
	err := o.Read(&user, "Username")

	if err == orm.ErrNoRows {
		c.Data["error"] = "รหัสผู้ใช้/รหัสผ่านไม่ถูกต้อง"
	} else {
		if errCript := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordForm)); errCript != nil {
			c.Data["error"] = "รหัสผู้ใช้/รหัสผ่านไม่ถูกต้อง"
		} else {
			if ok, errLogin := h.KeepLogin(c.Ctx.ResponseWriter, usernameForm, ""); ok == true {
				c.Ctx.Redirect(http.StatusFound, "/")
			} else {
				c.Data["error"] = errLogin
			}
		}
	}
	c.TplName = "auth/index.html"
	c.Render()
}

//Get to logout
func (c *LogoutController) Get() {
	h.LogOut(c.Ctx.ResponseWriter)
	c.Ctx.Redirect(http.StatusFound, "/auth")
}
