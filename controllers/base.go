package controllers

import (
	h "fix2man/helps"
	"fmt"

	"github.com/astaxie/beego"
)

//BaseController Login Validate
type BaseController struct {
	beego.Controller
}

//Prepare Login Validate
func (b *BaseController) Prepare() {
	fmt.Println(b.Ctx.Request.URL)
	val := h.GetUser(b.Ctx.Request)
	if val == "" {
		b.Ctx.Redirect(302, "/auth")
	}

	//Todo กำหนด รหัส Menu แล้วใส่ เป็น Data
	b.Data["active_p_001"] = "active menu-open"
	b.Data["active_c_001"] = "active"
}
