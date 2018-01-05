package controllers

import (
	h "fix2man/helps"
	m "fix2man/models"
	"fmt"
	"html/template"
	"net/url"
	"strings"

	"github.com/go-playground/form"
)

//RecController _
type RecController struct {
	BaseController
}

func parseForm() url.Values {
	return url.Values{
		"Name":                []string{"joeybloggs"},
		"Age":                 []string{"3"},
		"Gender":              []string{"Male"},
		"Address[0].Name":     []string{"26 Here Blvd."},
		"Address[0].Phone":    []string{"9(999)999-9999"},
		"Address[1].Name":     []string{"26 There Blvd."},
		"Address[1].Phone":    []string{"1(111)111-1111"},
		"active":              []string{"true"},
		"MapExample[key]":     []string{"value"},
		"NestedMap[key][key]": []string{"value"},
		"NestedArray[0][0]":   []string{"value"},
	}
}

//Product _
type Product struct {
	Name string
	Code string
}
type ReqX struct {
	Product []Product
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
	fmt.Println(input)
	var user ReqX
	err := decoder.Decode(&user, input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.Product[0].Code)
	fmt.Printf("%#v\n", user)
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
