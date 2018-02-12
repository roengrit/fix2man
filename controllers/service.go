package controllers

import (
	h "fix2man/helps"
	m "fix2man/models"
	"html/template"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//EmptyDateString _
func EmptyDateString(in string) (out string) {
	if in == "01-01-0001" {
		out = ""
	} else {
		out = in
	}
	return
}

//ServiceController _
type ServiceController struct {
	BaseController
}

//ServiceNonAuthController _
type ServiceNonAuthController struct {
	beego.Controller
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

//GetTechListJSON  _
func (c *ServiceController) GetTechListJSON() {
	term := strings.TrimSpace(c.GetString("query"))
	rowCount, err, lists := m.GetTechList("15", term)
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

//CalItemAvg  _
func (c *ServiceNonAuthController) CalItemAvg() {
	ret := m.NormalModel{}
	m.CalAllAvg()
	c.Data["json"] = ret
	c.ServeJSON()
}

//CalItemAvgByID _
func (c *ServiceNonAuthController) CalItemAvgByID() {
	ret := m.NormalModel{}
	ID, _ := strconv.ParseInt(c.GetString("id"), 10, 32)
	if ID != 0 {

		o := orm.NewOrm()
		_, _ = o.Raw("insert into stock_adj(product_id) values(?)  ", ID).Exec()
		m.CalAllAvg()
	}
	c.Data["json"] = ret
	c.ServeJSON()
}
