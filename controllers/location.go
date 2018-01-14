package controllers

import (
	h "fix2man/helps"
	m "fix2man/models"
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/form"
)

//LocationController _
type LocationController struct {
	BaseController
}

//DepartList _
func (c *LocationController) DepartList() {
	c.Data["title"] = "แผนก"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "location/depart-list.html"
	c.Data["branch"] = m.GetAllBranch()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "location/depart-list-script.html"
	c.Render()
}

//GetDepartList _
func (c *LocationController) GetDepartList() {
	ret := m.NormalModel{}
	ret.RetOK = true
	top, _ := strconv.ParseInt(c.GetString("top"), 10, 32)
	branchID := c.GetString("txt-branch")
	term := c.GetString("txt-search")
	lists, rowCount, err := m.GetDepartList(term, branchID, int(top))
	if err == nil {
		ret.RetOK = true
		ret.RetCount = int64(rowCount)
		ret.RetData = h.GenDepartHTML(*lists)
		if rowCount == 0 {
			ret.RetData = h.HTMLDepartNotFoundRows
		}
	} else {
		ret.RetOK = false
		ret.RetData = strings.Replace(h.HTMLDepartError, "{err}", err.Error(), -1)
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//CreateDepart _
func (c *LocationController) CreateDepart() {
	departID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	c.Data["title"] = "สร้าง/แก้ไขแผนก"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if departID != 0 {
		ret, _ := m.GetDepartByID(int(departID))
		c.Data["data"] = ret
	}
	c.Layout = "layout.html"
	c.TplName = "location/depart.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "location/depart-script.html"
	c.Render()
}

//GetDepart _
func (c *LocationController) GetDepart() {
	departID, _ := c.GetInt("id")
	c.Data["title"] = "สร้าง/แก้ไขแผนก"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if departID != 0 {
		ret, _ := m.GetDepartByID(departID)
		c.Data["data"] = ret
	}
	c.Layout = "layout.html"
	c.TplName = "location/depart.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "location/depart-script.html"
	c.Render()
}

//UpdateDepart _
func (c *LocationController) UpdateDepart() {
	var depart m.Departs
	decoder := form.NewDecoder()
	err := decoder.Decode(&depart, c.Ctx.Request.Form)
	ret := m.NormalModel{}
	ret.RetOK = true

	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else if depart.Name == "" {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุชื่อ"
	}

	if depart.Branch == nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุสาขา"
	}
	if ret.RetOK && depart.ID == 0 {
		depart.CreatedAt = time.Now()
		_, err := m.CreateDeparts(depart)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	} else if ret.RetOK && depart.ID > 0 {
		depart.UpdatedAt = time.Now()
		err := m.UpdateDeparts(depart)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//DeleteDepart _
func (c *LocationController) DeleteDepart() {
	ID, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	ret := m.NormalModel{}
	ret.RetOK = true
	err := m.DeleteDepartByID(int(ID))
	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else {
		ret.RetData = "ลบข้อมูลสำเร็จ"
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

/////////////////////////////////////////////////// อาคาร ///////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////

//BuildingList _
func (c *LocationController) BuildingList() {
	c.Data["title"] = "อาคาร"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "location/building-list.html"
	c.Data["branch"] = m.GetAllBranch()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "location/building-list-script.html"
	c.Render()
}

//GetBuildingList _
func (c *LocationController) GetBuildingList() {
	ret := m.NormalModel{}
	ret.RetOK = true
	top, _ := strconv.ParseInt(c.GetString("top"), 10, 32)
	branchID := c.GetString("txt-branch")
	term := c.GetString("txt-search")
	lists, rowCount, err := m.GetBuildingsList(term, branchID, int(top))
	if err == nil {
		ret.RetOK = true
		ret.RetCount = int64(rowCount)
		ret.RetData = h.GenBuildingHTML(*lists)
		if rowCount == 0 {
			ret.RetData = h.HTMLBuildingNotFoundRows
		}
	} else {
		ret.RetOK = false
		ret.RetData = strings.Replace(h.HTMLBuildingError, "{err}", err.Error(), -1)
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//CreateBuilding _
func (c *LocationController) CreateBuilding() {
	buildingID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	c.Data["title"] = "สร้าง/แก้ไขอาคาร"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if buildingID != 0 {
		ret, _ := m.GetBuildingsByID(int(buildingID))
		c.Data["data"] = ret
	}
	c.Layout = "layout.html"
	c.TplName = "location/building.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "location/building-script.html"
	c.Render()
}

//GetBuilding _
func (c *LocationController) GetBuilding() {
	buildingID, _ := c.GetInt("id")
	c.Data["title"] = "สร้าง/แก้ไขแผนก"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if buildingID != 0 {
		ret, _ := m.GetBuildingsByID(buildingID)
		c.Data["data"] = ret
	}
	c.Layout = "layout.html"
	c.TplName = "location/building.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "location/building-script.html"
	c.Render()
}

//UpdateBuilding _
func (c *LocationController) UpdateBuilding() {
	var building m.Buildings
	decoder := form.NewDecoder()
	err := decoder.Decode(&building, c.Ctx.Request.Form)
	ret := m.NormalModel{}
	ret.RetOK = true

	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else if building.Name == "" {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุชื่อ"
	}

	if building.Branch == nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุสาขา"
	}
	if ret.RetOK && building.ID == 0 {
		building.CreatedAt = time.Now()
		_, err := m.CreateBuildings(building)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	} else if ret.RetOK && building.ID > 0 {
		building.UpdatedAt = time.Now()
		err := m.UpdateBuildings(building)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//DeleteBuilding _
func (c *LocationController) DeleteBuilding() {
	ID, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	ret := m.NormalModel{}
	ret.RetOK = true
	err := m.DeleteBuildingsByID(int(ID))
	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else {
		ret.RetData = "ลบข้อมูลสำเร็จ"
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

/////////////////////////////////////////////////// ชั้น ///////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////

//ClassList _
func (c *LocationController) ClassList() {
	c.Data["title"] = "ชั้น"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "location/class-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "location/class-list-script.html"
	c.Render()
}

//GetClassList _
func (c *LocationController) GetClassList() {
	ret := m.NormalModel{}
	ret.RetOK = true
	top, _ := strconv.ParseInt(c.GetString("top"), 10, 32)
	branchID := c.GetString("Branch.ID")
	buildingID := c.GetString("Building.ID")
	term := c.GetString("txt-search")
	lists, rowCount, err := m.GetClassList(term, branchID, buildingID, int(top))
	if err == nil {
		ret.RetOK = true
		ret.RetCount = int64(rowCount)
		ret.RetData = h.GenClassHTML(*lists)
		if rowCount == 0 {
			ret.RetData = h.HTMLClassNotFoundRows
		}
	} else {
		ret.RetOK = false
		ret.RetData = strings.Replace(h.HTMLClassError, "{err}", err.Error(), -1)
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//CreateClass _
func (c *LocationController) CreateClass() {
	classID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	c.Data["title"] = "สร้าง/แก้ไขชั้น"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if classID != 0 {
		ret, _ := m.GetClassByID(int(classID))
		c.Data["data"] = ret
	}
	c.Layout = "layout.html"
	c.TplName = "location/class.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "location/class-script.html"
	c.Render()
}

//UpdateClass _
func (c *LocationController) UpdateClass() {
	var class m.Class
	decoder := form.NewDecoder()
	err := decoder.Decode(&class, c.Ctx.Request.Form)
	ret := m.NormalModel{}
	ret.RetOK = true

	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	}

	if c.GetString("Building.Branch.ID") == "" && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุสาขา"
	}

	if c.GetString("Building.ID") == "" && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุอาคาร"
	}

	if class.Building != nil && ret.RetOK {
		if class.Building.Branch == nil {
			ret.RetOK = false
			ret.RetData = "กรุณาระบุสาขา"
		}
	}

	if class.Building == nil && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุอาคาร"
	}
	if class.Building != nil && ret.RetOK {
		if class.Building.ID == 0 {
			ret.RetOK = false
			ret.RetData = "กรุณาระบุอาคาร"
		}
	}

	if class.Name == "" && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุชื่อ"
	}

	if ret.RetOK && class.ID == 0 {
		class.CreatedAt = time.Now()
		_, err := m.CreateClass(class)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	} else if ret.RetOK && class.ID > 0 {
		class.UpdatedAt = time.Now()
		err := m.UpdateClass(class)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//DeleteClass _
func (c *LocationController) DeleteClass() {
	ID, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	ret := m.NormalModel{}
	ret.RetOK = true
	err := m.DeleteClassByID(int(ID))
	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else {
		ret.RetData = "ลบข้อมูลสำเร็จ"
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}
