package controllers

import (
	"errors"
	h "fix2man/helps"
	m "fix2man/models"
	"html/template"
	"strconv"
	"strings"
	"time"
)

//ReqController _
type ReqController struct {
	BaseController
}

//Get _
func (c *ReqController) Get() {
	id := c.Ctx.Request.URL.Query().Get("id")
	now := time.Now()
	c.Data["title"] = "สร้างใบแจ้งงาน"
	c.Data["retCount"] = "0"
	c.Data["currentDate"] = now.Format("2006-01-02")
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if id != "" {
		ret, err := m.GetReqDocID(id)
		if err != nil {
			c.Data["err"] = err.Error()
		} else {
			c.Data["data"] = ret
			if nil != ret.User {
				c.Data["userID"] = ret.User.ID
			} else {
				c.Data["userID"] = ""
			}
			c.Data["currentDate"] = ret.EventDate.Format("2006-01-02")
		}
	}

	c.Layout = "layout.html"
	c.TplName = "req/req-read.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "req/req-style.tpl"
	c.LayoutSections["Scripts"] = "req/req-script.tpl"
	c.Render()
}

//Read _
func (c *ReqController) Read() {
	id := c.Ctx.Request.URL.Query().Get("id")
	now := time.Now()
	c.Data["title"] = "สร้างใบแจ้งงาน"
	c.Data["retCount"] = "0"
	c.Data["currentDate"] = now.Format("2006-01-02")
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["r"] = "readonly"
	if id != "" {
		ret, err := m.GetReqDocID(id)
		if err != nil {
			c.Data["err"] = err.Error()
		} else {
			c.Data["data"] = ret
			if nil != ret.User {
				c.Data["userID"] = ret.User.ID
			} else {
				c.Data["userID"] = ""
			}
			c.Data["currentDate"] = ret.EventDate.Format("2006-01-02")
		}
	}

	c.Layout = "layout.html"
	c.TplName = "req/req.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "req/req-style.tpl"
	c.LayoutSections["Scripts"] = "req/req-read-script.tpl"
	c.Render()
}

//Post _
func (c *ReqController) Post() {

	reqDocID, _ := strconv.ParseInt(c.GetString("req-doc-id"), 10, 32)
	reqNameID, _ := strconv.ParseInt(c.GetString("req-name-id"), 10, 32)
	reqName := c.GetString("req-name")
	reqBranchID, _ := strconv.ParseInt(c.GetString("req-branch-id"), 10, 32)
	reqDepartID, _ := strconv.ParseInt(c.GetString("req-depart-id"), 10, 32)
	reqBuildingID, _ := strconv.ParseInt(c.GetString("req-building-id"), 10, 32)
	reqClassID, _ := strconv.ParseInt(c.GetString("req-class-id"), 10, 32)
	reqRoomID, _ := strconv.ParseInt(c.GetString("req-room-id"), 10, 32)
	reqSn := c.GetString("req-sn")
	reqDateEvent := c.GetString("req-date-event")
	reqTel := c.GetString("req-tel")
	remark := c.GetString("remark")
	actionUser, _ := m.GetUserByUserName(h.GetUser(c.Ctx.Request))
	ret := m.NormalModel{}
	ret.RetOK = true

	if reqNameID == 0 && reqName == "" && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุชื่อผู้แจ้ง"
	}
	if reqTel == "" && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุเบอร์โทรศัพท์"
	}
	if reqBranchID == 0 && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุสาขา"
	}
	if reqDepartID == 0 && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาแผนก"
	}
	if reqBuildingID == 0 && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาตึก"
	}
	if reqClassID == 0 && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาชั้น"
	}
	if reqRoomID == 0 && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "กรุณาห้อง"
	}
	if reqDateEvent == "" && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "วันที่เสีย"
	}
	errDate := errors.New("")
	timeStamp := time.Now()
	sp := strings.Split(reqDateEvent, "-")
	if len(sp) == 3 {
		timeStamp, errDate = time.Parse("2006-02-01", sp[2]+"-"+sp[0]+"-"+sp[1])
	}
	if errDate != nil {
		ret.RetOK = false
		ret.RetData = "วันที่เสียไม่ถูกต้อง (dd-mm-yyyy)"
	}

	if remark == "" && ret.RetOK {
		ret.RetOK = false
		ret.RetData = "รายละเอียดการชำรุด/ปัญหา/อาการเสีย"
	}

	if ret.RetOK {
		newReq := m.RequestDocument{}
		newReq.Branch = &m.Branchs{ID: int(reqBranchID)}
		newReq.User = &m.Users{ID: int(reqNameID)}
		newReq.Building = &m.Buildings{ID: int(reqBuildingID)}
		newReq.Class = &m.Class{ID: int(reqClassID)}
		newReq.Room = &m.Rooms{ID: int(reqRoomID)}
		newReq.Depart = &m.Departs{ID: int(reqDepartID)}
		newReq.ReqName = reqName
		newReq.Tel = reqTel
		newReq.EventDate = timeStamp
		newReq.SerailNumber = reqSn
		newReq.Details = remark

		errAction := errors.New("")
		if reqDocID == 0 {
			newReq.DocDate = time.Now()
			newReq.DocNo = m.GetMaxDoc("request_document", "REQ")
			newReq.CreateUser = actionUser
			newReq.CreatedAt = time.Now()
			ret.ID, errAction = m.CreateReq(newReq, m.Users{ID: actionUser.ID})
		} else {
			newReq.ID = int(reqDocID)
			newReq.UpdatedAt = time.Now()
			newReq.UpdateUser = actionUser
			errAction = m.UpdateReq(newReq)
		}
		if errAction == nil {
			if reqDocID == 0 {
				ret.RetData = newReq.DocNo
				ret.FlagAction = "add"
			} else {
				ret.FlagAction = "edit"
				ret.RetData = "บันทึกสำเร็จ"
			}
			ret.XSRF = c.XSRFToken()
		} else {
			ret.RetOK = false
			ret.RetData = errAction.Error()
		}
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//ReqList _
func (c *ReqController) ReqList() {
	c.Data["title"] = "รายการใบแจ้งงาน"
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

//GetReqList _
func (c *ReqController) GetReqList() {

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
		ret.RetData = h.GenReqHTML(lists)
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