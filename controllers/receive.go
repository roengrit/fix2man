package controllers

import (
	"bytes"
	h "fix2man/helps"
	m "fix2man/models"
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/form"
)

//ReceiveController _
type ReceiveController struct {
	BaseController
}

//Get _
func (c *ReceiveController) Get() {
	docID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	if strings.Contains(c.Ctx.Request.URL.RequestURI(), "read") {
		c.Data["r"] = "readonly"
	}
	if docID == 0 {
		c.Data["title"] = "รับสินค้า/วัตถุดิบ"
	} else {
		doc, _ := m.GetReceive(int(docID))
		c.Data["m"] = doc
		if !doc.Active {
			c.Data["r"] = "readonly"
		}
		c.Data["RetCount"] = len(doc.ReceiveSub)
		c.Data["title"] = "แก้ไข รับสินค้า/วัตถุดิบ : " + doc.DocNo
	}
	c.Data["CurrentDate"] = time.Now()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "receive/receive.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["html_head"] = "receive/receive-style.html"
	c.LayoutSections["scripts"] = "receive/receive-script.html"
	c.Render()
}

//Post _
func (c *ReceiveController) Post() {
	doc := m.Receive{}
	doc.Flag = 1 // รับ
	actionUser, _ := m.GetUser(h.GetUser(c.Ctx.Request))
	retJSON := m.NormalModel{RetOK: true}
	decoder := form.NewDecoder()
	parsFormErr := decoder.Decode(&doc, c.Ctx.Request.Form)
	if parsFormErr == nil {
		if docDate, err := h.ValidateDate(c.GetString("DocDate")); err == nil {
			doc.DocDate = docDate
		} else {
			retJSON.RetOK = false
			retJSON.RetData = "มีข้อมูลบางอย่างไม่ครบถ้วน"
		}
		if retJSON.RetOK && doc.ID == 0 {
			_, parsFormErr = m.CreateReceive(doc, actionUser)
			if parsFormErr == nil {
				retJSON.RetOK = true
				retJSON.RetData = "บันทึกสำเร็จ"
			} else {
				retJSON.RetOK = false
				retJSON.RetData = parsFormErr.Error()
			}
		}
		if retJSON.RetOK && doc.ID != 0 {
			_, parsFormErr = m.UpdateReceive(doc, actionUser)
			if parsFormErr == nil {
				retJSON.RetOK = true
				retJSON.RetData = "บันทึกสำเร็จ"
			} else {
				retJSON.RetOK = false
				retJSON.RetData = parsFormErr.Error()
			}
		}
	} else {
		retJSON.RetOK = false
		retJSON.RetData = "มีข้อมูลบางอย่างไม่ครบถ้วน"
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = retJSON
	c.ServeJSON()
}

//ReceiveList _
func (c *ReceiveController) ReceiveList() {
	c.Data["beginDate"] = time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02")
	c.Data["endDate"] = time.Date(time.Now().Year(), time.Now().Month()+1, 0, 0, 0, 0, 0, time.Local).Format("2006-01-02")
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["title"] = "รับสินค้า/วัตถุดิบ"
	c.Layout = "layout.html"
	c.TplName = "receive/receive-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["html_head"] = "receive/receive-style.html"
	c.LayoutSections["scripts"] = "receive/receive-list-script.html"
	c.Render()
}

//GetReceiveList _
func (c *ReceiveController) GetReceiveList() {
	ret := m.NormalModel{}
	ret.RetOK = true
	top, _ := strconv.ParseInt(c.GetString("top"), 10, 32)
	term := c.GetString("txt-search")
	dateBegin := c.GetString("txt-date-begin")
	dateEnd := c.GetString("txt-date-end")
	if dateBegin != "" {
		sp := strings.Split(dateBegin, "-")
		dateBegin = sp[2] + "-" + sp[1] + "-" + sp[0]
	}
	if dateEnd != "" {
		sp := strings.Split(dateEnd, "-")
		dateEnd = sp[2] + "-" + sp[1] + "-" + sp[0]
	}
	lists, rowCount, err := m.GetReceiveList(term, int(top), dateBegin, dateEnd)
	if err == nil {
		ret.RetOK = true
		ret.RetCount = int64(rowCount)
		ret.RetData = h.GenReceiveHTML(*lists)
		if rowCount == 0 {
			ret.RetData = h.HTMLReceiveNotFoundRows
		}
	} else {
		ret.RetOK = false
		ret.RetData = strings.Replace(h.HTMLReceiveError, "{err}", err.Error(), -1)
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//CancelReceive _
func (c *ReceiveController) CancelReceive() {
	ID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	ret := m.NormalModel{}
	dataTemplate := m.NormalModel{}
	dataTemplate.ID = ID
	dataTemplate.Title = "กรุณาระบุ หมายเหตุ การยกเลิก"
	dataTemplate.XSRF = c.XSRFToken()
	t, err := template.ParseFiles("views/receive/receive-cancel.html")
	var tpl bytes.Buffer
	if err = t.Execute(&tpl, dataTemplate); err != nil {
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

//UpdateCancelReceive _
func (c *ReceiveController) UpdateCancelReceive() {
	actionUser, _ := m.GetUser(h.GetUser(c.Ctx.Request))
	ret := m.NormalModel{}
	ID, _ := c.GetInt("ID")
	remark := c.GetString("Remark")
	ret.RetOK = true
	if remark == "" {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุหมายเหตุ"
	}
	if ret.RetOK {
		_, err := m.UpdateCancelReceive(ID, remark, actionUser)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		}
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}
