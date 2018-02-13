package models

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//RequestDocument _
type RequestDocument struct {
	ID           int
	DocNo        string `orm:"size(20)"`
	DocType      int
	DocDate      time.Time
	ReqName      string     `orm:"size(255)"`
	User         *Users     `orm:"null;rel(fk)"`
	Tel          string     `orm:"size(50)"`
	Branch       *Branchs   `orm:"null;rel(one)"`
	Depart       *Departs   `orm:"null;rel(one)"`
	Building     *Buildings `orm:"null;rel(one)"`
	Class        *Class     `orm:"null;rel(one)"`
	Room         *Rooms     `orm:"null;rel(one)"`
	Location     string     `orm:"size(225)"`
	SerailNumber string     `orm:"size(50)"`

	EventDate       time.Time `form:"-"orm:"null"`
	EventTime       string    `orm:"null"`
	ReqDate         time.Time `form:"-"orm:"null"`
	ReqTime         string    `orm:"null"`
	AppointmentDate time.Time `form:"-"orm:"null"`
	AppointmentTime string    `orm:"null"`
	GoalDate        time.Time `form:"-"orm:"null"`
	GoalTime        string    `orm:"null"`
	ActionDate      time.Time `form:"-"orm:"null"`
	ActionTime      string    `orm:"null"`
	CompleteDate    time.Time `form:"-"orm:"null"`
	CompleteTime    string    `orm:"null"`

	EstimatePrice float64
	OtherPrice    float64
	TotalWork     float64
	TotalMaterial float64
	TotalAmount   float64

	TimeDiff float64

	Details  string `orm:"size(500)"`
	Remark   string `orm:"size(300)"`
	DocRefNo string `orm:"size(50)"`

	ActionNumber int
	ActionUser   []RequestUserAction `orm:"-"`
	CreateUser   *Users              `orm:"rel(one)"`
	CreatedAt    time.Time           `orm:"auto_now_add"`
	UpdateUser   *Users              `orm:"null;rel(one)"`
	UpdatedAt    time.Time           `orm:"null"`
}

//RequestList _
type RequestList struct {
	ID        int
	DocNo     string
	ReqDate   time.Time
	DocDate   time.Time
	ReqName   string
	Tel       string
	Branch    string
	Details   string
	EventDate time.Time
	Status    string
}

//RequestStatus _
type RequestStatus struct {
	ID              int
	RequestDocument *RequestDocument `orm:"rel(one)"`
	Status          *Status          `orm:"rel(one)"`
	Remark          string           `orm:"size(300)"`
	CreateUser      *Users           `orm:"rel(one)"`
	CreatedAt       time.Time        `orm:"auto_now_add"`
}

//RequestUserAction _
type RequestUserAction struct {
	ID              int
	Cost            float64          `orm:"digits(12);decimals(2)"`
	RequestDocument *RequestDocument `orm:"rel(one)"`
	ActionUser      *Users           `orm:"rel(one)"`
	CreatedAt       time.Time        `orm:"auto_now_add"`
}

//Status _
type Status struct {
	ID        int
	Name      string `orm:"size(225)"`
	IsDef     bool
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//DocRef _
type DocRef struct {
	ID          int
	Name        string `orm:"size(225)"`
	DocDate     time.Time
	DocNo       string
	TotalAmount float64
}

func init() {
	orm.RegisterModel(new(RequestDocument), new(RequestStatus), new(Status), new(RequestUserAction)) // Need to register model in init
}

//GetAllStatus _
func GetAllStatus() (req *[]Status) {
	o := orm.NewOrm()
	reqGet := &[]Status{}
	o.QueryTable("status").RelatedSel().OrderBy("ID").All(reqGet)
	return reqGet
}

//GetAllStatusExcludeID _
func GetAllStatusExcludeID(ID int) (req *[]Status) {
	o := orm.NewOrm()
	reqGet := &[]Status{}
	o.QueryTable("status").RelatedSel().OrderBy("ID").Exclude("ID__in", ID).All(reqGet)
	return reqGet
}

//GetFirstStatus _
func GetFirstStatus() (req *Status) {
	o := orm.NewOrm()
	reqGet := &Status{}
	o.QueryTable("status").RelatedSel().OrderBy("ID").Filter("is_def", true).One(reqGet)
	return reqGet
}

//CreateReq _
func CreateReq(req RequestDocument, user Users) (retID int64, retDocNo string, errRet error) {
	req.DocNo = GetMaxDoc("request_document", "REQ")
	req.DocDate = time.Now()
	req.CreatedAt = time.Now()
	req.CreateUser = &user
	o := orm.NewOrm()
	var firstStatus = GetFirstStatus()
	firstInsertStatus := RequestStatus{RequestDocument: &req, Status: firstStatus, CreateUser: &user, CreatedAt: time.Now()}
	errRet = o.Begin()
	if errRet == nil {
		id, err := o.Insert(&req)
		userActions := []RequestUserAction{}
		for _, val := range req.ActionUser {
			val.RequestDocument = &RequestDocument{ID: int(id)}
			if val.ActionUser.ID != 0 {
				valUser, _ := GetUserByUserID(val.ActionUser.ID)
				val.Cost = valUser.CostPerTechnical
				userActions = append(userActions, val)
			}
		}
		_, err = o.Insert(&firstInsertStatus)
		if len(userActions) > 0 {
			_, err = o.InsertMulti(len(userActions), userActions)
		}
		if err != nil {
			errRet = err
			err = o.Rollback()
		} else {
			o.Commit()
			UpdateReqTotalAmount(req.DocNo)
		}
		errRet = err
	}
	return retID, req.DocNo, errRet
}

//UpdateReq _
func UpdateReq(req RequestDocument, user Users) (errRet error) {
	o := orm.NewOrm()
	doc := RequestDocument{ID: req.ID}
	if o.Read(&doc) == nil {
		doc.Branch = req.Branch
		doc.Building = req.Building
		doc.Class = req.Class
		doc.Depart = req.Depart
		doc.Details = req.Details
		doc.EventDate = req.EventDate
		doc.EventTime = req.EventTime
		doc.ReqDate = req.ReqDate
		doc.ReqTime = req.ReqTime
		doc.TimeDiff = req.TimeDiff
		doc.AppointmentDate = req.AppointmentDate
		doc.AppointmentTime = req.AppointmentTime
		doc.GoalDate = req.GoalDate
		doc.GoalTime = req.GoalTime
		doc.ActionDate = req.ActionDate
		doc.ActionTime = req.ActionTime
		doc.CompleteDate = req.CompleteDate
		doc.CompleteTime = req.CompleteTime

		doc.ActionNumber = req.ActionNumber
		doc.EstimatePrice = req.EstimatePrice
		doc.OtherPrice = req.OtherPrice
		doc.Remark = req.Remark
		doc.DocRefNo = req.DocRefNo
		doc.ReqName = req.ReqName
		doc.Room = req.Room
		doc.SerailNumber = req.SerailNumber
		doc.Tel = req.Tel
		doc.UpdatedAt = req.UpdatedAt
		doc.UpdateUser = req.UpdateUser
		doc.User = req.User
		errRet = o.Begin()
		if errRet == nil {
			if _, err := o.Update(&doc); err == nil {
				_, err = o.QueryTable("request_user_action").Filter("request_document_id", doc.ID).Delete()
				userActions := []RequestUserAction{}
				for _, val := range req.ActionUser {
					val.RequestDocument = &doc
					if val.ActionUser.ID != 0 {
						valUser, _ := GetUserByUserID(val.ActionUser.ID)
						val.Cost = valUser.CostPerTechnical
						userActions = append(userActions, val)
					}
				}
				if len(userActions) > 0 {
					_, err = o.InsertMulti(len(userActions), userActions)
				}
				if err != nil {
					errRet = err
					err = o.Rollback()
				} else {
					o.Commit()
					UpdateReqTotalAmount(doc.DocNo)
				}
			} else {
				errRet = err
				err = o.Rollback()
			}
		}
	}
	return errRet
}

//UpdateReqTotalAmount _
func UpdateReqTotalAmount(docNo string) {
	o := orm.NewOrm()
	req := &RequestDocument{}
	o.QueryTable("request_document").Filter("DocNo", docNo).RelatedSel().One(req)
	if req.ID == 0 {
		return
	}
	reqActionUser, _ := GetReqUserActionByDocID(strconv.Itoa(req.ID))
	reqDocRef, _ := GetDocRef(docNo)
	var Total, TotalWork, TotalMaterial float64
	Total = req.OtherPrice
	for _, val := range *reqActionUser {
		Total = Total + (val.Cost * req.TimeDiff)
		TotalWork = TotalWork + (val.Cost * req.TimeDiff)
	}
	for _, val := range reqDocRef {
		Total = Total + val.TotalAmount
		TotalMaterial = TotalMaterial + val.TotalAmount
	}
	o.Raw("update request_document  set total_amount = ? , total_work = ? , total_material = ?  where i_d = ?", Total, TotalWork, TotalMaterial, req.ID).Exec()
}

//GetReqDocID _
func GetReqDocID(ID string) (req *RequestDocument, errRet error) {
	o := orm.NewOrm()
	id, _ := strconv.Atoi(ID)
	reqGet := &RequestDocument{}
	o.QueryTable("request_document").Filter("ID", id).RelatedSel().One(reqGet)
	if nil != reqGet {
		if nil != reqGet.User {
			reqGet.User.Password = ""
		}
		reqGet.CreateUser = nil
		reqGet.UpdateUser = nil
		actionUser, _ := GetReqUserActionByDocID(ID)
		reqGet.ActionUser = *actionUser
	}
	if reqGet.ID == 0 {
		return nil, errors.New("ไม่พบข้อมูล")
	}
	return reqGet, errRet
}

//GetReqUserActionByDocID _
func GetReqUserActionByDocID(ID string) (req *[]RequestUserAction, errRet error) {
	o := orm.NewOrm()
	reqGet := &[]RequestUserAction{}
	o.QueryTable("request_user_action").Filter("request_document_id", ID).RelatedSel().All(reqGet)
	return reqGet, errRet
}

//GetReqDocLastStatus _
func GetReqDocLastStatus(ID int) (req *RequestStatus, errRet error) {
	o := orm.NewOrm()
	reqGet := &RequestStatus{}
	o.QueryTable("request_status").Filter("request_document_id", ID).RelatedSel().OrderBy("-created_at").One(reqGet)
	if reqGet.ID == 0 {
		return nil, errors.New("ไม่พบข้อมูล")
	}
	return reqGet, errRet
}

//GetReqDocStatusList _
func GetReqDocStatusList(ID int) (req *[]RequestStatus, errRet error) {
	o := orm.NewOrm()
	reqGet := &[]RequestStatus{}
	o.QueryTable("request_status").Filter("request_document_id", ID).RelatedSel().OrderBy("created_at").All(reqGet)
	return reqGet, errRet
}

//GetReqDocHasStatus _
func GetReqDocHasStatus(docID, statusID int) (req *RequestStatus, errRet error) {
	o := orm.NewOrm()
	reqGet := &RequestStatus{}
	o.QueryTable("request_status").Filter("request_document_id", docID).Filter("status_id", statusID).RelatedSel().OrderBy("-created_at").One(reqGet)
	if reqGet.ID == 0 {
		return nil, errors.New("ไม่พบข้อมูล")
	}
	return reqGet, errRet
}

//CreateReqStatus _
func CreateReqStatus(reqStatus RequestStatus) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&reqStatus)
	o.Commit()
	if err == nil {
		retID = id
	}
	return retID, err
}

//GetDocRef _
func GetDocRef(docNo string) (docList []DocRef, errRet error) {
	o := orm.NewOrm()
	sql := `SELECT
				pick_up.i_d,
				doc_no,
				doc_date,
				total_amount,
				users.name
			FROM
				pick_up
				JOIN users ON pick_up.creator_id = users.i_d
			WHERE 	pick_up.active and pick_up.doc_ref_no = ?`
	_, err := o.Raw(sql, docNo).QueryRows(&docList)
	return docList, err
}

//GetReqDocList _
func GetReqDocList(top, term, branch, status, dateBegin, dateEnd string) (num int64, err error, reqList []RequestList) {

	var sql = `SELECT  i_d,doc_no,doc_date,req_date,req_name,tel,
						(
							SELECT NAME FROM Branchs WHERE i_d = branch_id  
						) as branch,
						details,
						event_date,
					    (
							SELECT ( SELECT NAME FROM STATUS WHERE STATUS.i_d = st.status_id ) 
							FROM request_status st 
							WHERE st.request_document_id = request_document.i_d 
							ORDER BY st.created_at DESC 
							LIMIT 1 
						)  as status
			   FROM request_document    
			   WHERE (lower(doc_no) like lower(?) or lower(req_name) like lower(?) or lower(details) like lower(?)) and (1=1)
			  `
	var filterSQL = ""
	if branch != "" {
		if filterSQL != "" {
			filterSQL += " AND branch_id = " + branch
		} else {
			filterSQL += " branch_id = " + branch
		}
	}
	if status != "" {
		if filterSQL != "" {
			filterSQL += ` AND (SELECT st.status_id  
								FROM request_status st  
								WHERE st.request_document_id = request_document.i_d  
								ORDER BY st.created_at 
								DESC  LIMIT 1  ) = ` + status
		} else {
			filterSQL += ` (SELECT st.status_id  
							FROM request_status st  
							WHERE st.request_document_id = request_document.i_d  
							ORDER BY st.created_at DESC  
							LIMIT 1  ) = ` + status
		}
	}
	if dateBegin != "" && dateEnd != "" {
		if filterSQL != "" {
			filterSQL += ` AND DATE(doc_date) between '` + dateBegin + `' AND '` + dateEnd + `'`
		} else {
			filterSQL += ` DATE(doc_date) between '` + dateBegin + `' AND '` + dateEnd + `'`
		}
	}
	if filterSQL != "" {
		sql = strings.Replace(sql, "1=1", filterSQL, -1)
	}
	sql = sql + ` ORDER BY doc_no desc LIMIT {0}`
	if top == "0" {
		sql = strings.Replace(sql, "LIMIT {0}", "", -1)
	} else {
		sql = strings.Replace(sql, "{0}", top, -1)
	}
	o := orm.NewOrm()
	num, err = o.Raw(sql, "%"+term+"%", "%"+term+"%", "%"+term+"%").QueryRows(&reqList)
	return num, err, reqList
}
