package controllers

import (
	m "fix2man/models"
	"html/template"
	"strconv"
	"time"

	"github.com/go-playground/form"
)

//SupplierController _
type SupplierController struct {
	BaseController
}

//SupList _
func (c *SupplierController) SupList() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["title"] = "ร้านค้า/Supplier"
	c.Layout = "layout.html"
	c.TplName = "supplier/sup-list.html"
	c.Render()
}

//GetSupList _
func (c *SupplierController) GetSupList() {
	ret := m.NormalModel{}
	ret.RetOK = true
	top, _ := strconv.ParseInt(c.GetString("top"), 10, 32)
	term := c.GetString("txt-search")
	retSup, _ := m.GetSuppliersList(term, int(top))
	ret.Data1 = retSup
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//CreateSup _
func (c *SupplierController) CreateSup() {
	supID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	c.Data["Province"] = m.GetAllProvince()
	if supID == 0 {
		c.Data["title"] = "สร้าง ร้านค้า/Supplier"
	} else {
		c.Data["title"] = "แก้ไข ร้านค้า/Supplier"
		sup, _ := m.GetSuppliers(int(supID))
		c.Data["data"] = sup
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "supplier/sup.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "supplier/sup-script.html"
	c.Render()
}

//UpdateSup _
func (c *SupplierController) UpdateSup() {
	c.Data["title"] = "หน้าหลัก"
	var sub m.Suppliers
	decoder := form.NewDecoder()
	err := decoder.Decode(&sub, c.Ctx.Request.Form)
	ret := m.NormalModel{}
	ret.RetOK = true
	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else if c.GetString("Name") == "" {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุชื่อ"
	}

	if ret.RetOK && sub.ID == 0 {
		sub.CreatedAt = time.Now()
		_, err := m.CreateSuppliers(sub)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	} else if ret.RetOK && sub.ID > 0 {
		sub.UpdatedAt = time.Now()
		err := m.UpdateSuppliers(sub)
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
