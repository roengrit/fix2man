package controllers

import (
	h "fix2man/helps"
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
	case s.Contains(uri, "user") || s.Contains(uri, "role"):
		{
			b.Data["m_user"] = "active menu-open"
			if s.Contains(uri, "/normal/?entity=roles") {
				b.Data["m_role_list"] = "active"
			}
		}
	case s.Contains(uri, "entity") || s.Contains(uri, "normal"):
		{
			b.Data["m_setting"] = "active menu-open"
			if s.Contains(uri, "/normal/?entity=status") {
				b.Data["m_status_list"] = "active"
			}
			if s.Contains(uri, "/normal/?entity=categorys") {
				b.Data["m_category_list"] = "active"
			}
			if s.Contains(uri, "/normal/?entity=units") {
				b.Data["m_unit_list"] = "active"
			}
			if s.Contains(uri, "/normal/?entity=branchs") {
				b.Data["m_branch_list"] = "active"
			}
			if s.Contains(uri, "/entity/location/depart") {
				b.Data["m_depart_list"] = "active"
			}
		}
	}

}
