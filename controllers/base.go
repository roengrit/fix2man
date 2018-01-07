package controllers

import (
	h "fix2man/helps"
	"fmt"
	s "strings"

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
	b.Data["UserDisplay"] = val
	//Todo กำหนด รหัส Menu แล้วใส่ เป็น Data
	uri := b.Ctx.Request.URL.RequestURI()
	fmt.Println(uri)
	switch {
	case s.Contains(uri, "request"):
		{
			b.Data["m_request"] = "active menu-open"
			if s.Contains(uri, "/create-request") {
				b.Data["m_create_request"] = "active"
			}
			if s.Contains(uri, "/request/list") {
				b.Data["m_request_list"] = "active"
			}
		}
	case s.Contains(uri, "supplier"):
		{
			b.Data["m_supplier"] = "active menu-open"
			if s.Contains(uri, "/supplier/list") {
				b.Data["m_supplier_list"] = "active"
			}
		}
	}

}
