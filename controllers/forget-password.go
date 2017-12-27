package controllers

import (
	h "fix2man/helps"
	m "fix2man/models"
	"html/template"

	"github.com/astaxie/beego"
)

//ForgetController _
type ForgetController struct {
	beego.Controller
}

//Get _
func (c *ForgetController) Get() {

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["username"] = ""
	c.Data["title"] = "กรอก email เพื่อรับรหัสผ่าน"
	c.TplName = "forget-password/forget.html"
	c.Render()
}

//Post to
func (c *ForgetController) Post() {
	usernameForm := c.GetString("username")
	newPass := m.RandStringRunes(8)
	if hasUser, errFindeuser := m.GetUser(usernameForm); hasUser {
		if errSendMail := h.SendMail(usernameForm, newPass); errSendMail == "" {
			if ok, err := m.ForgetPass(usernameForm, newPass); ok {
				c.Data["success"] = "ส่งรหัสผ่านสำเร็จ"
			} else {
				c.Data["error"] = err
			}
		} else {
			c.Data["error"] = errSendMail
		}
	} else {
		c.Data["error"] = errFindeuser
	}

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["username"] = c.GetString("username")

	c.TplName = "forget-password/forget.html"
	c.Render()
}
