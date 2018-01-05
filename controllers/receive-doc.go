package controllers

import (
	h "fix2man/helps"
	m "fix2man/models"
	"fmt"
	"html/template"
	"net/url"
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
	Name string
	Code string
}

//Suplier _
type Suplier struct {
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
	c.LayoutSections["HtmlHead"] = "receive/rec-style.html"
	c.LayoutSections["Scripts"] = "receive/rec-script.html"
	c.Render()
}

//Post _
func (c *RecController) Post() {
	c.Data["title"] = "สร้างใบรับ"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "receive/rec.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "receive/rec-style.html"
	c.LayoutSections["Scripts"] = "receive/rec-script.html"
	decoder := form.NewDecoder()
	c.Ctx.Request.ParseForm()
	var input url.Values
	input = c.Ctx.Request.Form
	var user RecDocument
	err := decoder.Decode(&user, input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.Suplier.ID)
	fmt.Println(user.Suplier.Name)
	fmt.Println(user.Project.Name)
	fmt.Println(user.Product[0].Code)
	c.Render()
}

//RecList _
func (c *RecController) RecList() {
	c.Data["title"] = "รายการใบรับ"
	c.Layout = "layout.html"
	c.TplName = "receive/rec-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "receive/rec-style.html"
	c.LayoutSections["Scripts"] = "receive/rec-list-script.html"
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
