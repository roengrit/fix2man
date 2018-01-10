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
	EventDate    time.Time  `form:"-"orm:"null"`
	ReqDate      time.Time  `form:"-"orm:"null"`
	Details      string     `orm:"size(500)"`
	Remark       string     `orm:"size(300)"`
	DocRefNo     string     `orm:"size(50)"`
	CreateUser   *Users     `orm:"rel(one)"`
	CreatedAt    time.Time  `orm:"auto_now_add"`
	UpdateUser   *Users     `orm:"null;rel(one)"`
	UpdatedAt    time.Time  `orm:"null"`
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

//Status _
type Status struct {
	ID        int
	Name      string `orm:"size(225)"`
	IsDef     bool
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	orm.RegisterModel(new(RequestDocument), new(RequestStatus), new(Status)) // Need to register model in init
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
	first_status := RequestStatus{RequestDocument: &req, Status: firstStatus, CreateUser: &user, CreatedAt: time.Now()}
	o.Begin()
	id, err := o.Insert(&req)
	_, err = o.Insert(&first_status)
	o.Commit()
	if err == nil {
		retID = id
	}
	return retID, req.DocNo, err
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
		doc.ReqDate = req.ReqDate
		doc.DocRefNo = req.DocRefNo
		doc.ReqName = req.ReqName
		doc.Room = req.Room
		doc.SerailNumber = req.SerailNumber
		doc.Tel = req.Tel
		doc.UpdatedAt = req.UpdatedAt
		doc.UpdateUser = req.UpdateUser
		doc.User = req.User
		if num, err := o.Update(&doc); err == nil {
			_ = num
		} else {
			errRet = err
		}
	}
	return errRet
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
	}
	if reqGet.ID == 0 {
		return nil, errors.New("ไม่พบข้อมูล")
	}
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
