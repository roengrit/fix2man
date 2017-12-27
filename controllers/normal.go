package controllers

import (
	"bytes"
	h "fix2man/helps"
	m "fix2man/models"
	"html/template"
)

//EntitryController _
type EntitryController struct {
	BaseController
}

//Get _
func (c *EntitryController) Get() {
	entity := c.Ctx.Request.URL.Query().Get("entity")
	title := h.GetEntityTitle(entity)
	c.Data["title"] = title
	c.Data["retCount"] = "0"
	c.Data["entity"] = entity
	c.Layout = "layout.html"
	c.TplName = "normal/normal.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "normal/normal-script.tpl"
	c.Render()
}

//ListEntity  _
func (c *EntitryController) ListEntity() {

	term := c.GetString("txt-search")
	top := c.GetString("top")
	entity := c.GetString("entity")

	num, err, lists := m.GetListEntity(entity, top, term)

	ret := m.NormalModel{}

	if err == nil {
		ret.RetOK = true
		ret.RetCount = num
		ret.RetData = h.GenEntityHtml(lists)
		if num == 0 {
			ret.RetData = `<tr><td></td><td>*** ไม่พบข้อมูล ***</td><td></td></tr>`
		}

	} else {
		ret.RetOK = false
		ret.RetData = `<tr><td></td><td>` + err.Error() + `</td><td></td></tr>`
	}

	c.Data["json"] = ret
	c.ServeJSON()
}

//NewEntity _
func (c *EntitryController) NewEntity() {
	entity := c.Ctx.Request.URL.Query().Get("entity")
	title := h.GetEntityTitle(entity)
	ret := m.NormalModel{}
	t, err := template.ParseFiles("views/normal/normal-add.html")
	max := m.GetMaxEntity(entity)
	var tpl bytes.Buffer
	if err = t.Execute(&tpl, map[string]string{"entity": entity, "title": title, "code": max, "id": ""}); err != nil {
		ret.RetOK = err != nil
		ret.RetData = err.Error()
	} else {
		ret.RetOK = true
		ret.RetData = tpl.String()
	}
	c.Data["json"] = ret
	c.ServeJSON()
}

//MaxEntity _
func (c *EntitryController) MaxEntity() {
	entity := c.Ctx.Request.URL.Query().Get("entity")
	ret := m.NormalModel{}
	max := m.GetMaxEntity(entity)
	ret.RetOK = true
	ret.RetData = max
	c.Data["json"] = ret
	c.ServeJSON()
}

//GetEntity _
func (c *EntitryController) GetEntity() {
	c.Data["json"] = ""
	c.ServeJSON()
}

//UpdateEntity _
func (c *EntitryController) UpdateEntity() {
	c.Data["json"] = ""
	c.ServeJSON()
}
