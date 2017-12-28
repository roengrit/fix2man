package controllers

import (
	c "fix2man/controllers"
	"html/template"
)

//RoleController _
type RoleController struct {
	c.BaseController
}

//Get _
func (c *RoleController) Get() {
	c.Data["title"] = "จัดการสิทธิ์"
	c.Data["retCount"] = "0"
	c.Data["entity"] = "roles"
	c.Layout = "layout.html"
	c.TplName = "role-member/role/role.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "role-member/role/role.tpl"
	c.Render()
}
