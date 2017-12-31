package controllers

import (
	"html/template"
	"time"
)

//ReqController _
type ReqController struct {
	BaseController
}

//Get _
func (c *ReqController) Get() {
	now := time.Now()
	c.Data["title"] = "สร้างใบแจ้งงาน"
	c.Data["retCount"] = "0"
	c.Data["currentDate"] = now.Format("02/01/2006")
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "req/req.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "req/req-style.tpl"
	c.LayoutSections["Scripts"] = "req/req-script.tpl"
	c.Render()
}
