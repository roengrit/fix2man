package controllers

import (
	m "fix2man/models"
	"html/template"
	"strconv"
	"time"

	"github.com/go-playground/form"
)

//LocationController _
type LocationController struct {
	BaseController
}

//GetDepartList _
func (c *LocationController) GetDepartList() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "location/depart-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "location/depart-list-script.html"
	c.Render()
}

//CreateDepart _
func (c *LocationController) CreateDepart() {
	departID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	c.Data["title"] = "สร้าง/แก้ไขแผนก"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if departID != 0 {
		ret, _ := m.GetDepartByID(int(departID))
		c.Data["data"] = ret
	}
	c.Layout = "layout.html"
	c.TplName = "location/depart.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "location/depart-script.html"
	c.Render()
}

//UpdateDepart _
func (c *LocationController) UpdateDepart() {
	var depart m.Departs
	decoder := form.NewDecoder()
	err := decoder.Decode(&depart, c.Ctx.Request.Form)
	ret := m.NormalModel{}
	ret.RetOK = true

	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else if depart.Name == "" {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุชื่อ"
	}

	if depart.Branch == nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุสาขา"
	}
	if ret.RetOK && depart.ID == 0 {
		depart.CreatedAt = time.Now()
		_, err := m.CreateDeparts(depart)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	} else if ret.RetOK && depart.ID > 0 {
		depart.UpdatedAt = time.Now()
		err := m.UpdateDeparts(depart)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

// //ListEntity  _
// func (c *LocationController) ListEntity() {

// 	term := c.GetString("txt-search")
// 	top := c.GetString("top")
// 	entity := c.GetString("entity")
// 	title := h.GetEntityTitle(entity)
// 	ret := m.NormalModel{}
// 	if title != "" {
// 		rowCount, err, lists := m.GetListEntity(entity, top, term)
// 		if err == nil {
// 			ret.RetOK = true
// 			ret.RetCount = rowCount
// 			ret.RetData = h.GenEntityHTML(lists)
// 			if rowCount == 0 {
// 				ret.RetData = h.HTMLNotFoundRows
// 			}
// 		} else {
// 			ret.RetOK = false
// 			ret.RetData = strings.Replace(h.HTMLError, "{err}", err.Error(), -1)
// 		}
// 	} else {
// 		ret.RetData = h.HTMLPermissionDenie
// 	}
// 	c.Data["json"] = ret
// 	c.ServeJSON()
// }

// //ListEntityJSON  _
// func (c *LocationController) ListEntityJSON() {

// 	term := c.GetString("query")
// 	entity := c.Ctx.Request.URL.Query().Get("entity")
// 	title := h.GetEntityTitle(entity)
// 	ret := m.NormalModel{}
// 	if title != "" {
// 		rowCount, err, lists := m.GetListEntity(entity, "15", term)
// 		if err == nil {
// 			ret.RetOK = true
// 			ret.RetCount = rowCount
// 			ret.ListData = lists
// 			if rowCount == 0 {
// 				ret.RetOK = false
// 				ret.RetData = "ไม่พบข้อมูล"
// 			}
// 		} else {
// 			ret.RetOK = false
// 			ret.RetData = "ไม่พบข้อมูล"
// 		}
// 	} else {
// 		ret.RetData = "ไม่พบข้อมูล"
// 	}
// 	c.Data["json"] = ret
// 	c.ServeJSON()
// }

// //NewEntity _
// func (c *LocationController) NewEntity() {
// 	entity := c.Ctx.Request.URL.Query().Get("entity")
// 	ID := c.Ctx.Request.URL.Query().Get("id")
// 	del := c.Ctx.Request.URL.Query().Get("del")
// 	title := h.GetEntityTitle(entity)

// 	ret := m.NormalModel{}
// 	var code, name, alert string

// 	if del != "" {
// 		title = "คุณต้องการลบข้อมูล " + title + " ใช่หรือไม่"
// 	}

// 	t, err := template.ParseFiles("views/normal/normal-add.html")

// 	if ID == "" {

// 	} else {
// 		errGet, retVal := m.GetEntity(entity, ID)

// 		if errGet == nil && (m.NormalEntity{}) != retVal {
// 			code = retVal.Code
// 			name = retVal.Name
// 		} else {
// 			alert = "ไม่พบข้อมูล"
// 		}
// 	}

// 	var tpl bytes.Buffer

// 	tplVal := map[string]string{
// 		"entity": entity, "title": title,
// 		"code": code, "id": ID, "name": name,
// 		"alert": alert, "del": del,
// 		"xsrfdata": c.XSRFToken()}

// 	if err = t.Execute(&tpl, tplVal); err != nil {
// 		ret.RetOK = err != nil
// 		ret.RetData = err.Error()
// 	} else {
// 		ret.RetOK = true
// 		ret.RetData = tpl.String()
// 	}

// 	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
// 	c.Data["json"] = ret
// 	c.ServeJSON()
// }

// //UpdateEntity _
// func (c *LocationController) UpdateEntity() {
// 	entity := c.GetString("entity")
// 	ID := c.GetString("narmal-id")
// 	name := c.GetString("normal-name")
// 	del := c.GetString("del-flag")
// 	title := h.GetEntityTitle(entity)

// 	ret := m.NormalModel{}
// 	ret.RetOK = true

// 	if title == "" {
// 		ret.RetData = "ไม่อนุญาติ ใน entity อื่น"
// 		ret.RetOK = false
// 	}

// 	if del != "" && ret.RetOK { // ลบ
// 		err := m.DeleteEntity(entity, ID)
// 		if err == nil {
// 			ret.RetOK = true
// 			ret.RetData = "ลบข้อมูลสำเร็จ"
// 		} else {
// 			ret.RetOK = false
// 			ret.RetData = err.Error()
// 		}
// 	}

// 	if ID != "" && del == "" && ret.RetOK { // แก้ไข
// 		err := m.UpdateEntity(entity, ID, name)
// 		if err == nil {
// 			ret.RetOK = true
// 			ret.RetData = "แก้ไขข้อมูลสำเร็จ"
// 		} else {
// 			ret.RetOK = false
// 			ret.RetData = err.Error()
// 		}
// 	}

// 	if ID == "" && del == "" && ret.RetOK { // สร้าง
// 		_, err := m.CreateEntity(entity, name)
// 		if err == nil {
// 			ret.RetOK = true
// 			ret.RetData = "บันทึกข้อมูลสำเร็จ"
// 		} else {
// 			ret.RetOK = false
// 			ret.RetData = err.Error()
// 		}
// 	}

// 	c.Data["json"] = ret
// 	c.ServeJSON()
// }
