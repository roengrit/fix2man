package controllers

import (
	"bytes"
	h "fix2man/helps"
	m "fix2man/models"
	"html/template"
	"strings"
)

//EntitryController _
type EntitryController struct {
	BaseController
}

//Get _
func (c *EntitryController) Get() {
	entity := c.Ctx.Request.URL.Query().Get("entity")
	title := h.GetEntityTitle(entity)
	if title == "" {
		c.Ctx.WriteString("*** ไม่อนุญาติ ใน entity อื่น ***")
	} else {
		c.Data["title"] = title
		c.Data["retCount"] = "0"
		c.Data["entity"] = entity
		c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
		c.Layout = "layout.html"
		c.TplName = "normal/normal.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "normal/normal-script.html"
		c.Render()
	}
}

//ListEntity  _
func (c *EntitryController) ListEntity() {

	term := c.GetString("txt-search")
	top := c.GetString("top")
	entity := c.GetString("entity")
	title := h.GetEntityTitle(entity)
	ret := m.NormalModel{}
	if title != "" {
		rowCount, err, lists := m.GetListEntity(entity, top, term)
		if err == nil {
			ret.RetOK = true
			ret.RetCount = rowCount
			ret.RetData = h.GenEntityHTML(lists)
			if rowCount == 0 {
				ret.RetData = h.HTMLNotFoundRows
			}
		} else {
			ret.RetOK = false
			ret.RetData = strings.Replace(h.HTMLError, "{err}", err.Error(), -1)
		}
	} else {
		ret.RetData = h.HTMLPermissionDenie
	}
	c.Data["json"] = ret
	c.ServeJSON()
}

//ListEntityJSON  _
func (c *EntitryController) ListEntityJSON() {

	term := c.GetString("query")
	entity := c.Ctx.Request.URL.Query().Get("entity")
	title := h.GetEntityTitle(entity)
	ret := m.NormalModel{}
	if title != "" {
		rowCount, err, lists := m.GetListEntity(entity, "15", term)
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
	} else {
		ret.RetData = "ไม่พบข้อมูล"
	}
	c.Data["json"] = ret
	c.ServeJSON()
}

//NewEntity _
func (c *EntitryController) NewEntity() {
	entity := c.Ctx.Request.URL.Query().Get("entity")
	id := c.Ctx.Request.URL.Query().Get("id")
	del := c.Ctx.Request.URL.Query().Get("del")
	title := h.GetEntityTitle(entity)

	ret := m.NormalModel{}
	var code, name, alert string

	if del != "" {
		title = "คุณต้องการลบข้อมูล " + title + " ใช่หรือไม่"
	}

	t, err := template.ParseFiles("views/normal/normal-add.html")

	if id == "" {

	} else {
		errGet, retVal := m.GetEntity(entity, id)

		if errGet == nil && (m.NormalEntity{}) != retVal {
			code = retVal.Code
			name = retVal.Name
		} else {
			alert = "ไม่พบข้อมูล"
		}
	}

	var tpl bytes.Buffer

	tplVal := map[string]string{
		"entity": entity, "title": title,
		"code": code, "id": id, "name": name,
		"alert": alert, "del": del,
		"xsrfdata": c.XSRFToken()}

	if err = t.Execute(&tpl, tplVal); err != nil {
		ret.RetOK = err != nil
		ret.RetData = err.Error()
	} else {
		ret.RetOK = true
		ret.RetData = tpl.String()
	}

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//UpdateEntity _
func (c *EntitryController) UpdateEntity() {
	entity := c.GetString("entity")
	id := c.GetString("narmal-id")
	name := c.GetString("normal-name")
	del := c.GetString("del-flag")
	title := h.GetEntityTitle(entity)

	ret := m.NormalModel{}
	ret.RetOK = true

	if title == "" {
		ret.RetData = "ไม่อนุญาติ ใน entity อื่น"
		ret.RetOK = false
	}

	if del != "" && ret.RetOK { // ลบ
		err := m.DeleteEntity(entity, id)
		if err == nil {
			ret.RetOK = true
			ret.RetData = "ลบข้อมูลสำเร็จ"
		} else {
			ret.RetOK = false
			ret.RetData = err.Error()
		}
	}

	if id != "" && del == "" && ret.RetOK { // แก้ไข
		err := m.UpdateEntity(entity, id, name)
		if err == nil {
			ret.RetOK = true
			ret.RetData = "แก้ไขข้อมูลสำเร็จ"
		} else {
			ret.RetOK = false
			ret.RetData = err.Error()
		}
	}

	if id == "" && del == "" && ret.RetOK { // สร้าง
		_, err := m.CreateEntity(entity, name)
		if err == nil {
			ret.RetOK = true
			ret.RetData = "บันทึกข้อมูลสำเร็จ"
		} else {
			ret.RetOK = false
			ret.RetData = err.Error()
		}
	}

	c.Data["json"] = ret
	c.ServeJSON()
}
