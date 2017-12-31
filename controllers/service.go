package controllers

import (
	h "fix2man/helps"
	m "fix2man/models"
	"fmt"
	"strings"
)

//ReqController _
type ServiceController struct {
	BaseController
}

//ListEntityJson  _
func (c *ServiceController) ListEntityJson() {

	term := strings.TrimSpace(c.GetString("query"))
	entity := c.Ctx.Request.URL.Query().Get("entity")
	ret := m.NormalModel{}
	rowCount, err, lists := m.GetListEntity(entity, "15", term)
	if err == nil {
		ret.RetOK = true
		ret.RetCount = rowCount
		ret.ListData = lists
		if rowCount == 0 {
			ret.RetOK = false
			ret.RetData = "ไม่พบข้อมูล"
		}
	} else {
		ret.RetOK = false
		ret.RetData = "ไม่พบข้อมูล"
	}
	c.Data["json"] = ret
	c.ServeJSON()
}

//ListEntityWithParentJson  _
func (c *ServiceController) ListEntityWithParentJson() {

	term := strings.TrimSpace(c.GetString("query"))
	entity := c.Ctx.Request.URL.Query().Get("entity")
	parent := strings.TrimSpace(c.GetString("parent"))
	parentEntity := h.GetEntityParentField(entity)
	ret := m.NormalModel{}
	fmt.Println(entity)
	fmt.Println(parentEntity)
	fmt.Println(parent)
	rowCount, err, lists := m.GetListEntityWithParent(entity, parentEntity, "15", parent, term)
	fmt.Println(lists)
	if err == nil {
		ret.RetOK = true
		ret.RetCount = rowCount
		ret.ListData = lists
		if rowCount == 0 {
			ret.RetOK = false
			ret.RetData = "ไม่พบข้อมูล"
		}
	} else {
		ret.RetOK = false
		ret.RetData = "ไม่พบข้อมูล"
	}
	c.Data["json"] = ret
	c.ServeJSON()
}

//GetUserJson  _
func (c *ServiceController) GetUserListJson() {
	term := strings.TrimSpace(c.GetString("query"))
	rowCount, err, lists := m.GetUserList("15", term)
	ret := m.NormalModel{}
	if err == nil {
		ret.RetOK = true
		ret.RetCount = rowCount
		ret.ListData = lists
		if rowCount == 0 {
			ret.RetOK = false
			ret.RetData = "ไม่พบข้อมูล"
		}
	} else {
		ret.RetOK = false
		ret.RetData = "ไม่พบข้อมูล"
	}
	c.Data["json"] = ret
	c.ServeJSON()
}

//GetBuildingListJson  _
// func (c *ServiceController) GetBuildingListJson() {
// 	term := strings.TrimSpace(c.GetString("query"))
// 	branch := strings.TrimSpace(c.GetString("branch"))
// 	rowCount, err, lists := m.GetBuildingList("15", branch, term)
// 	ret := m.NormalModel{}
// 	if err == nil {
// 		ret.RetOK = true
// 		ret.RetCount = rowCount
// 		ret.ListData = lists
// 		if rowCount == 0 {
// 			ret.RetOK = false
// 			ret.RetData = "ไม่พบข้อมูล"
// 		}
// 	} else {
// 		ret.RetOK = false
// 		ret.RetData = "ไม่พบข้อมูล"
// 	}
// 	c.Data["json"] = ret
// 	c.ServeJSON()
// }

//GetUserJson  _
func (c *ServiceController) GetUserJson() {
	term := strings.TrimSpace(c.GetString("query"))
	user, err := m.GetUserByID(term)
	if err != "" {
		c.Data["json"] = err
	} else {
		c.Data["json"] = user
	}
	c.ServeJSON()
}
