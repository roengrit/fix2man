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

func (a *EntitryController) Prepare() {
	a.EnableXSRF = false
}

//Get _
func (c *EntitryController) Get() {
	val := h.GetUser(c.Ctx.Request)
	if val == "" {
		c.Ctx.Redirect(302, "/auth")
	}
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
	title := h.GetEntityTitle(entity)
	ret := m.NormalModel{}
	if title != "" {
		num, err, lists := m.GetListEntity(entity, top, term)
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
	} else {
		ret.RetData = `<tr><td></td><td>*** ไม่อนุญาติ ใน entity อื่น ***</td><td></td></tr>`
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
	if del != "" {
		title = "คุณต้องการลบข้อมูล " + title + " ใช่หรือไม่"
	}
	ret := m.NormalModel{}
	t, err := template.ParseFiles("views/normal/normal-add.html")
	var code, name, alert string
	if id == "" {
		code = m.GetMaxEntity(entity)
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
	if err = t.Execute(&tpl, map[string]string{"entity": entity, "title": title, "code": code, "id": id, "name": name, "alert": alert, "del": del}); err != nil {
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
	entity := c.Ctx.Request.URL.Query().Get("entity")
	id := c.GetString("narmal-id")
	title := h.GetEntityTitle(entity)
	ret := m.NormalModel{}
	if title != "" {
		err, retVal := m.GetEntity(entity, id)
		if err != nil {
			ret.RetOK = true
			ret.ID = int64(retVal.ID)
			ret.Name = retVal.Name
		} else {
			ret.RetData = err.Error()
		}
	} else {
		ret.RetData = "ไม่อนุญาติ ใน entity อื่น"
	}
	c.Data["json"] = ret
	c.ServeJSON()
}

//UpdateEntity _
func (c *EntitryController) UpdateEntity() {
	entity := c.GetString("entity")
	id := c.GetString("narmal-id")
	code := c.GetString("normal-code")
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
		err := m.UpdateEntity(entity, id, code, name)
		if err == nil {
			ret.RetOK = true
			ret.RetData = "แก้ไขข้อมูลสำเร็จ"
		} else {
			ret.RetOK = false
			ret.RetData = err.Error()
		}
	}

	if id == "" && del == "" && ret.RetOK { // สร้าง
		_, err := m.CreateEntity(entity, code, name)
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
