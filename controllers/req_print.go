package controllers

import (
	m "fix2man/models"
	"html/template"
)

//ReqPrintController _
type ReqPrintController struct {
	BaseController
}

//Get _
func (c *ReqPrintController) Get() {
	docRef := c.Ctx.Request.URL.Query().Get("doc_ref")
	isPrint := c.Ctx.Request.URL.Query().Get("print")
	c.Data["title"] = "ใบงานเลขที่ : " + docRef
	c.Data["docRef"] = docRef
	c.Data["isPrint"] = isPrint
	doc, _ := m.GetReqDocByDocNo(docRef)
	c.Data["m"] = doc
	statusList, _ := m.GetReqDocStatusList(int(doc.ID))
	docrefList, _ := m.GetDocRef(doc.DocNo)
	c.Data["user_len"] = len(doc.ActionUser)
	c.Data["status"] = statusList
	c.Data["doc_ref_list"] = docrefList
	c.Data["topic"] = m.GetAllTopic(docRef)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "req-print/index.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "req-print/index-script.html"
	c.Render()
}
