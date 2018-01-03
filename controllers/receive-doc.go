package controllers

import (
	h "fix2man/helps"
	m "fix2man/models"
	"html/template"
	"strings"
	"time"
)

//RecController _
type RecController struct {
	BaseController
}

//Get _
func (c *RecController) Get() {
	now := time.Now()
	c.Data["title"] = "สร้างใบรับสินค้า"
	c.Data["currentDate"] = now.Format("2006-01-02")
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "req/rec.html"
	c.LayoutSections = make(map[string]string)
	//c.LayoutSections["HtmlHead"] = "req/req-style.tpl"
	//c.LayoutSections["Scripts"] = "req/req-script.tpl"
	c.Render()
}

//RecList _
func (c *RecController) RecList() {
	c.Data["title"] = "รายการใบรับสินค้า"
	c.Data["beginDate"] = time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02")
	c.Data["endDate"] = time.Date(time.Now().Year(), time.Now().Month()+1, 0, 0, 0, 0, 0, time.Local).Format("2006-01-02")
	c.Data["retCount"] = "0"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["status"] = m.GetAllStatus()
	c.Data["branch"] = m.GetAllBranch()
	c.Layout = "layout.html"
	c.TplName = "req/req-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "req/req-style.tpl"
	c.LayoutSections["Scripts"] = "req/req-list-script.tpl"
	c.Render()
}

//GetRecList _
func (c *RecController) GetRecList() {

	top := c.GetString("top")
	term := c.GetString("txt-search")
	branch := c.GetString("txt-branch")
	status := c.GetString("txt-status")
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

	ret := m.NormalModel{}
	rowCount, err, lists := m.GetReqDocList(top, term, branch, status, dateBegin, dateEnd)
	if err == nil {
		ret.RetOK = true
		ret.RetCount = rowCount
		_ = lists
		ret.RetData = h.GenRecHTML(lists)
		if rowCount == 0 {
			ret.RetData = h.HTMLReqNotFoundRows
		}
	} else {
		ret.RetOK = false
		ret.RetData = strings.Replace(h.HTMLError, "{err}", err.Error(), -1)
	}

	c.Data["json"] = ret
	c.ServeJSON()
}
