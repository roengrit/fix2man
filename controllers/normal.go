package controllers

import (
	"bytes"
	m "fix2man/models"
	"html/template"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

//EntitryController _
type EntitryController struct {
	BaseController
}

//Get _
func (c *EntitryController) Get() {
	entity := c.Ctx.Request.URL.Query().Get("entity")
	switch entity {
	case "roles":
		c.Data["title"] = "จัดการสิทธิ์"
	case "units":
		c.Data["title"] = "จัดการหน่วย"
	case "status":
		c.Data["title"] = "สถานะการซ่อม"
	case "branchs":
		c.Data["title"] = "สาขา/ไซต์"
	default:
		return
	}

	c.Data["retCount"] = "0"
	c.Data["entity"] = entity
	c.Layout = "layout.html"
	c.TplName = "normal/normal.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "normal/normal-script.tpl"
	c.Render()
}

//ListRole _
func (c *EntitryController) ListEntity() {
	term := c.GetString("txt-search")
	top := c.GetString("top")
	entity := c.GetString("entity")
	var sql = "SELECT i_d,code, name FROM " + entity + " WHERE name like ? or code like ? limit {0}"
	if top == "0" {
		sql = strings.Replace(sql, "limit {0}", "", -1)
	} else {
		sql = strings.Replace(sql, "{0}", top, -1)
	}

	o := orm.NewOrm()
	var roles []m.Roles
	num, err := o.Raw(sql, "%"+term+"%", "%"+term+"%").QueryRows(&roles)

	ret := m.NormalModel{}

	var hmtlBuffer bytes.Buffer
	var htmlTemplate = `<tr>
								<td>
									{code} 
								</td>
								<td>{name}</td>
								<td>
										<div class="btn-group">
											<button type="button" class="btn btn-sm btn-primary" onclick='editNormal({id})'>แก้ไข</button>
											<button type="button" class="btn btn-sm btn-danger" onclick='deleteNormal({id})'>ลบ</button>
										</div>
								</td>                             
							</tr>`

	if err == nil {
		ret.RetOK = true
		for _, val := range roles {
			temp := strings.Replace(htmlTemplate, "{code}", val.Code, -1)
			temp = strings.Replace(temp, "{name}", val.Name, -1)
			temp = strings.Replace(temp, "{id}", strconv.Itoa(val.ID), -1)
			hmtlBuffer.WriteString(temp)
		}
		ret.RetCount = num
		ret.RetData = hmtlBuffer.String()
		if num == 0 {
			ret.RetData = `<tr><td></td><td>*** ไม่พบข้อมูล ***</td><td></td></tr>`
		}

	} else {
		ret.RetOK = false
		ret.RetData = `<tr><td></td><td>` + err.Error() + `</td><td></td></tr>`
	}

	c.Data["json"] = ret
	c.ServeJSON()
}

//NewRole _
func (c *EntitryController) NewEntity() {

	c.Data["json"] = ""
	c.ServeJSON()
}

func (c *EntitryController) GetEntity() {
	c.Data["title"] = "จัดการสิทธิ์"
	c.Data["retCount"] = "0"
	c.Layout = "layout.html"
	c.TplName = "role-member/role/role.html"
	c.LayoutSections["Scripts"] = "role-member/role/role.tpjs"
	c.Render()
}

func (c *EntitryController) UpdateEntity() {
	c.Data["title"] = "จัดการสิทธิ์"
	c.Data["retCount"] = "0"
	c.Layout = "layout.html"
	c.TplName = "role-member/role/role.html"
	c.LayoutSections["Scripts"] = "role-member/role/role.tpjs"
	c.Render()
}
