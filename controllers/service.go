package controllers

import (
	h "fix2man/helps"
	m "fix2man/models"
	"html/template"
	"strings"
)

//ServiceController _
type ServiceController struct {
	BaseController
}

//ListEntityJSON  _
func (c *ServiceController) ListEntityJSON() {

	term := strings.TrimSpace(c.GetString("query"))
	entity := c.Ctx.Request.URL.Query().Get("entity")
	ret := m.NormalModel{}
	rowCount, err, lists := m.GetListEntity(entity, 15, term)
	if err == nil {
		ret.RetOK = true
		ret.RetCount = rowCount
		ret.ListData = lists
		if rowCount == 0 {
			ret.RetOK = false
			ret.RetData = "ไม่พบข้อมูล"
		}
	} else {
		ret.RetOK = false
		ret.RetData = "ไม่พบข้อมูล"
	}
	c.Data["json"] = ret
	c.ServeJSON()
}

//ListEntityWithParentJSON  _
func (c *ServiceController) ListEntityWithParentJSON() {

	term := strings.TrimSpace(c.GetString("query"))
	entity := c.Ctx.Request.URL.Query().Get("entity")
	parent := strings.TrimSpace(c.GetString("parent"))
	parentEntity := h.GetEntityParentField(entity)
	ret := m.NormalModel{}
	rowCount, err, lists := m.GetListEntityWithParent(entity, parentEntity, 15, parent, term)
	if err == nil {
		ret.RetOK = true
		ret.RetCount = rowCount
		ret.ListData = lists
		if rowCount == 0 {
			ret.RetOK = false
			ret.RetData = "ไม่พบข้อมูล"
		}
	} else {
		ret.RetOK = false
		ret.RetData = "ไม่พบข้อมูล"
	}
	c.Data["json"] = ret
	c.ServeJSON()
}

//GetUserListJSON  _
func (c *ServiceController) GetUserListJSON() {
	term := strings.TrimSpace(c.GetString("query"))
	rowCount, err, lists := m.GetUserList("15", term)
	ret := m.NormalModel{}
	if err == nil {
		ret.RetOK = true
		ret.RetCount = rowCount
		ret.ListData = lists
		if rowCount == 0 {
			ret.RetOK = false
			ret.RetData = "ไม่พบข้อมูล"
		}
	} else {
		ret.RetOK = false
		ret.RetData = "ไม่พบข้อมูล"
	}
	c.Data["json"] = ret
	c.ServeJSON()
}

//GetUserJSON  _
func (c *ServiceController) GetUserJSON() {
	term := strings.TrimSpace(c.GetString("query"))
	user, err := m.GetUserByID(term)
	if err != "" {
		c.Data["json"] = err
	} else {
		c.Data["json"] = user
	}
	c.ServeJSON()
}

//GetXSRF _
func (c *ServiceController) GetXSRF() {
	c.Data["json"] = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.ServeJSON()
}
