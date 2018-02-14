package controllers

import (
	"html/template"
)

//AssessController _
type AssessController struct {
	BaseController
}

//Get _
func (c *AssessController) Get() {
	docRef := c.Ctx.Request.URL.Query().Get("doc_no")
	c.Data["title"] = "ประเมินการดำเนินงาน : " + docRef
	c.Data["docRef"] = docRef
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "assessment/index.html"
	c.LayoutSections = make(map[string]string)
	// c.LayoutSections["html_head"] = "pickup/pickup-style.html"
	c.LayoutSections["scripts"] = "assessment/index-script.html"
	c.Render()
}
