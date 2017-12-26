package controllers

import (
	h "fix2man/helps"

	"github.com/astaxie/beego"
)

//BaseController Login Validate
type BaseController struct {
	beego.Controller
}

//Prepare Login Validate
func (b *BaseController) Prepare() {
	val := h.GetUser(b.Ctx.Request)
	if val == "" {
		b.Ctx.Redirect(302, "/auth")
	}
}
