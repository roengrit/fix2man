package controllers

import (
	"bytes"
	"errors"
	h "fix2man/helps"
	m "fix2man/models"
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/form"
)

//ReqController _
type ReqController struct {
	BaseController
}

//Get _
func (c *ReqController) Get() {
	ID := c.Ctx.Request.URL.Query().Get("id")
	c.Data["req_doc_ref"] = c.Ctx.Request.URL.Query().Get("doc_ref")
	c.Data["title"] = "สร้างใบแจ้งงาน"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	if ID != "" {
		ret, err := m.GetReqDocID(ID)
		if err != nil {
			c.Data["err"] = err.Error()
		} else {
			c.Data["title"] = "แก้ไขใบแจ้งงาน"
			c.Data["data"] = ret
			c.Data["user_len"] = len(ret.ActionUser)
		}
	}
	if h.CheckPermissAllow(1001, c.Ctx.Request) {
		c.Data["can_change_action_number"] = "readonly"
	}
	c.Data["current_date"] = time.Now()
	c.Layout = "layout.html"
	c.TplName = "req/req.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["html_head"] = "req/req-style.html"
	c.LayoutSections["scripts"] = "req/req-script.html"
	c.Render()
}

//ReadReq _
func (c *ReqController) ReadReq() {
	ID := c.Ctx.Request.URL.Query().Get("id")
	docID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	c.Data["title"] = "ใบแจ้งงาน"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["r"] = "readonly"
	if ID != "" {
		ret, err := m.GetReqDocID(ID)
		if err != nil {
			c.Data["err"] = err.Error()
		} else {
			statusList, _ := m.GetReqDocStatusList(int(docID))
			c.Data["data"] = ret
			c.Data["status"] = statusList
		}
	}
	c.Layout = "layout.html"
	c.TplName = "req/req-read.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["html_head"] = "req/req-style.html"
	c.LayoutSections["scripts"] = "req/req-read-script.html"
	c.Render()
}

//Post _
func (c *ReqController) Post() {
	var reqDoc m.RequestDocument
	retJSON := m.NormalModel{}
	decoder := form.NewDecoder()
	parsFormErr := decoder.Decode(&reqDoc, c.Ctx.Request.Form)
	if parsFormErr == nil {
		reqEventDateString := c.GetString("EventDate")
		reqDateString := c.GetString("ReqDate")
		reqAppointmentDateString := c.GetString("AppointmentDate")
		reqGoalDateString := c.GetString("GoalDate")
		reqActionDateString := c.GetString("ActionDate")
		reqCompleteDateString := c.GetString("CompleteDate")
		actionUser, _ := m.GetUserByUserName(h.GetUser(c.Ctx.Request))
		ret, reqDateEvent, reqDate, reqAppointmentDate, reqGoalDate, reqActionDate, reqCompleteDate := h.ValidateReqData(reqDoc,
			reqDateString, reqEventDateString, reqAppointmentDateString,
			reqGoalDateString, reqActionDateString, reqCompleteDateString)
		if ret.RetOK {
			reqDoc.EventDate = reqDateEvent
			reqDoc.ReqDate = reqDate
			if reqAppointmentDate != (time.Time{}) {
				reqDoc.AppointmentDate = reqAppointmentDate
			}
			if reqGoalDate != (time.Time{}) {
				reqDoc.GoalDate = reqGoalDate
			}
			if reqActionDate != (time.Time{}) {
				reqDoc.ActionDate = reqActionDate
			}
			if reqCompleteDate != (time.Time{}) {
				reqDoc.CompleteDate = reqCompleteDate
			}
			errAction := errors.New("")
			if reqDoc.ID == 0 {
				ret.ID, reqDoc.DocNo, errAction = m.CreateReq(reqDoc, m.Users{ID: actionUser.ID})
			} else {
				errAction = m.UpdateReq(reqDoc, m.Users{ID: actionUser.ID})
			}
			if errAction == nil {
				if reqDoc.ID == 0 {
					ret.RetData = reqDoc.DocNo
				} else {
					ret.RetData = "บันทึกสำเร็จ"
				}
			} else {
				ret.RetOK = false
				ret.RetData = errAction.Error()
			}
		}
		retJSON = ret
	} else {
		retJSON.RetOK = false
		retJSON.RetData = "ข้อมูลผิดพลาด"
	}
	retJSON.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = retJSON
	c.ServeJSON()
}

//ReqList _
func (c *ReqController) ReqList() {
	c.Data["title"] = "รายการใบแจ้งงาน"
	c.Data["beginDate"] = time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02")
	c.Data["endDate"] = time.Date(time.Now().Year(), time.Now().Month()+1, 0, 0, 0, 0, 0, time.Local).Format("2006-01-02")
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["status"] = m.GetAllStatus()
	c.Data["branch"] = m.GetAllBranch()
	c.Layout = "layout.html"
	c.TplName = "req/req-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["html_head"] = "req/req-list-style.html"
	c.LayoutSections["scripts"] = "req/req-list-script.html"
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

//ChangeStatus _
func (c *ReqController) ChangeStatus() {
	ID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	ret := m.NormalModel{}
	lastStatus, _ := m.GetReqDocLastStatus(int(ID))
	dataTemplate := m.NormalModel{}
	dataTemplate.ID = ID
	dataTemplate.ListData = m.GetAllStatusExcludeID(lastStatus.Status.ID)
	dataTemplate.XSRF = c.XSRFToken()
	t, err := template.ParseFiles("views/req/req-change-status.html")
	var tpl bytes.Buffer

	if err = t.Execute(&tpl, dataTemplate); err != nil {
		ret.RetOK = err != nil
		ret.RetData = err.Error()
	} else {
		ret.RetOK = true
		ret.RetData = tpl.String()
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//UpdateStatus _
func (c *ReqController) UpdateStatus() {
	ret := m.NormalModel{}
	ID := c.GetString("req-id")
	docID, _ := strconv.ParseInt(ID, 10, 32)
	statusVal := c.GetString("txt-status")
	statusID, _ := strconv.ParseInt(statusVal, 10, 32)
	remark := c.GetString("remark")
	ret.RetOK = true
	if ret.RetOK {
		lastStatus, _ := m.GetReqDocLastStatus(int(docID))
		if int(statusID) == lastStatus.Status.ID {
			ret.RetOK = false
			ret.RetData = "เอกสารเป็นสถานะนี้อยู่แล้ว"
		}
	}
	if ret.RetOK && remark == "" {
		docHasStatus, errHas := m.GetReqDocHasStatus(int(docID), int(statusID))
		if errHas == nil && docHasStatus != nil && (&m.RequestStatus{}) != docHasStatus {
			if int(statusID) == docHasStatus.Status.ID {
				ret.RetOK = false
				ret.RetData = "คุณกำลังกลับไปสถานะก่อนหน้านี้ กรุณาระบุหมายเหตุ"
			}
		}
	}
	if ret.RetOK {
		actionUser, _ := m.GetUserByUserName(h.GetUser(c.Ctx.Request))
		status := m.RequestStatus{Remark: remark, CreateUser: actionUser, CreatedAt: time.Now()}
		status.Status = &m.Status{ID: int(statusID)}
		status.RequestDocument = &m.RequestDocument{ID: int(docID)}
		_, err := m.CreateReqStatus(status)
		if err != nil {
			ret.RetOK = false
			ret.RetData = "บันทึกไม่สำเร็จ " + err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//StartJob _
func (c *ReqController) StartJob() {
	ret := m.NormalModel{}
	ID := c.GetString("req-id")
	docID, _ := strconv.ParseInt(ID, 10, 32)
	statusVal := c.GetString("txt-status")
	statusID, _ := strconv.ParseInt(statusVal, 10, 32)
	remark := c.GetString("remark")
	ret.RetOK = true
	if ret.RetOK {
		lastStatus, _ := m.GetReqDocLastStatus(int(docID))
		if int(statusID) == lastStatus.Status.ID {
			ret.RetOK = false
			ret.RetData = "เอกสารเป็นสถานะนี้อยู่แล้ว"
		}
	}
	if ret.RetOK && remark == "" {
		docHasStatus, errHas := m.GetReqDocHasStatus(int(docID), int(statusID))
		if errHas == nil && docHasStatus != nil && (&m.RequestStatus{}) != docHasStatus {
			if int(statusID) == docHasStatus.Status.ID {
				ret.RetOK = false
				ret.RetData = "คุณกำลังกลับไปสถานะก่อนหน้านี้ กรุณาระบุหมายเหตุ"
			}
		}
	}
	if ret.RetOK {
		actionUser, _ := m.GetUserByUserName(h.GetUser(c.Ctx.Request))
		status := m.RequestStatus{Remark: remark, CreateUser: actionUser, CreatedAt: time.Now()}
		status.Status = &m.Status{ID: int(statusID)}
		status.RequestDocument = &m.RequestDocument{ID: int(docID)}
		_, err := m.CreateReqStatus(status)
		if err != nil {
			ret.RetOK = false
			ret.RetData = "บันทึกไม่สำเร็จ " + err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}
