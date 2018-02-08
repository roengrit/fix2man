package controllers

import (
	h "fix2man/helps"
	m "fix2man/models"
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/form"
)

//MemberController _
type MemberController struct {
	BaseController
}

//CreateMember _
func (c *MemberController) CreateMember() {
	memID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	if strings.Contains(c.Ctx.Request.URL.RequestURI(), "read") {
		c.Data["r"] = "readonly"
	}
	if memID == 0 {
		c.Data["title"] = "สร้าง สมาชิก"
	} else {
		c.Data["title"] = "แก้ไข สมาชิก"
		mem, _ := m.GetMember(int(memID))
		c.Data["m"] = mem
	}
	c.Data["Province"] = m.GetAllProvince()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "member/mem.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "member/mem-script.html"
	c.Render()
}

//UpdateMember _
func (c *MemberController) UpdateMember() {

	var sub m.Member
	decoder := form.NewDecoder()
	err := decoder.Decode(&sub, c.Ctx.Request.Form)
	ret := m.NormalModel{}
	actionUser, _ := m.GetUser(h.GetUser(c.Ctx.Request))

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
		sub.Creator = &actionUser
		_, err := m.CreateMember(sub)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	} else if ret.RetOK && sub.ID > 0 {
		sub.EditedAt = time.Now()
		sub.Editor = &actionUser
		err := m.UpdateMember(sub)
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

//DeleteMember _
func (c *MemberController) DeleteMember() {
	ID, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	ret := m.NormalModel{}
	ret.RetOK = true
	err := m.DeleteMember(int(ID))
	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else {
		ret.RetData = "ลบข้อมูลสำเร็จ"
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//MemberList _
func (c *MemberController) MemberList() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["title"] = "สมาชิก"
	c.Layout = "layout.html"
	c.TplName = "member/mem-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "member/mem-list-script.html"
	c.Render()
}

//GetMemberList _
func (c *MemberController) GetMemberList() {
	ret := m.NormalModel{}
	ret.RetOK = true
	top, _ := strconv.ParseInt(c.GetString("top"), 10, 32)
	term := c.GetString("txt-search")
	lists, rowCount, err := m.GetMemberList(term, int(top))
	if err == nil {
		ret.RetOK = true
		ret.RetCount = int64(rowCount)
		ret.RetData = h.GenMemHTML(*lists)
		if rowCount == 0 {
			ret.RetData = h.HTMLMemNotFoundRows
		}
	} else {
		ret.RetOK = false
		ret.RetData = strings.Replace(h.HTMLMemError, "{err}", err.Error(), -1)
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}
