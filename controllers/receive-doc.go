package controllers

import (
	h "fix2man/helps"
	m "fix2man/models"
	"strings"
)

//RecController _
type RecController struct {
	BaseController
}

//Get _
func (c *RecController) Get() {
	c.Data["title"] = "สร้างใบรับสินค้า"
	c.Layout = "layout.html"
	c.TplName = "receive/rec.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "receive/rec-list-style.tpl"
	c.LayoutSections["Scripts"] = "receive/rec-list-script.tpl"

	c.Render()
}

//RecList _
func (c *RecController) RecList() {
	c.Data["title"] = "รายการใบรับสินค้า"
	c.Layout = "layout.html"
	c.TplName = "receive/rec-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "receive/rec-style.tpl"
	c.LayoutSections["Scripts"] = "receive/rec-list-script.tpl"
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
