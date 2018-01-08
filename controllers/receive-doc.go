package controllers

import (
	h "fix2man/helps"
	m "fix2man/models"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/go-playground/form"
)

//RecController _
type RecController struct {
	BaseController
}

//Product _
type Product struct {
	ID   int
	Code string
	Name string
	Unit Unit
}

//Suplier _
type Suplier struct {
	ID   string
	Name string
}

//Unit _
type Unit struct {
	ID   string
	Name string
}

//Project _
type Project struct {
	ID   string
	Name string
}

//RecDocument _
type RecDocument struct {
	DocNo    string
	DocDate  time.Time
	DocRefNo string
	Remark   string
	Product  []Product
	Suplier  Suplier
	Project  Project
}

//Get _
func (c *RecController) Get() {
	c.Data["title"] = "สร้างใบรับ"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "receive/rec.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["html_head"] = "receive/rec-style.html"
	c.LayoutSections["scripts"] = "receive/rec-script.html"
	c.Render()
}

//Post _
func (c *RecController) Post() {
	c.Data["title"] = "สร้างใบรับ"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	decoder := form.NewDecoder()
	var recDoc RecDocument
	err := decoder.Decode(&recDoc, c.Ctx.Request.Form)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(recDoc.Suplier.ID)
	fmt.Println(recDoc.Suplier.Name)
	fmt.Println(recDoc.Project.ID)
	fmt.Println(recDoc.Project.Name)
	fmt.Println(recDoc.Product[0].Code)
	for _, ret := range recDoc.Product {
		fmt.Println(ret.Code)
	}
	//ret := m.NormalModel{}
	c.Data["json"] = recDoc
	c.ServeJSON()
}

//RecList _
func (c *RecController) RecList() {
	c.Data["title"] = "รายการใบรับ"
	c.Layout = "layout.html"
	c.TplName = "receive/rec-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["html_head"] = "receive/rec-style.html"
	c.LayoutSections["scripts"] = "receive/rec-list-script.html"
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
