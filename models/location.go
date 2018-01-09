package models

import (
	"errors"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Branchs _
type Branchs struct {
	ID                  int
	Name                string `orm:"size(225)"`
	Lock                bool
	TokenLine           string
	CostAvgPerTechnical float64
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

//Departs _
type Departs struct {
	ID        int
	Branch    *Branchs `orm:"rel(one)"`
	Name      string   `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time `orm:"auto_now_add"`
	UpdatedAt time.Time `orm:"null"`
}

//Buildings _
type Buildings struct {
	ID        int
	Branch    *Branchs `orm:"rel(fk)"`
	Name      string   `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Class _
type Class struct {
	ID        int
	Building  *Buildings `orm:"rel(one)"`
	Name      string     `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Rooms _
type Rooms struct {
	ID        int
	Class     *Class `orm:"rel(one)"`
	Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	orm.RegisterModel(
		new(Buildings),
		new(Class),
		new(Rooms),
		new(Branchs),
		new(Departs),
	) // Need to register model in init
}

//GetBuildingList _
func GetBuildingList(top, branchID, term string) (num int64, err error, buildList []Buildings) {
	var sql = "SELECT i_d,name FROM buildings WHERE branch_id = ? and lower(name) like lower(?) order by name limit {0}"
	if top == "0" {
		sql = strings.Replace(sql, "limit {0}", "", -1)
	} else {
		sql = strings.Replace(sql, "{0}", top, -1)
	}
	o := orm.NewOrm()
	num, err = o.Raw(sql, branchID, "%"+term+"%").QueryRows(&buildList)
	if err != nil {

	}
	return num, err, buildList
}

//GetAllBranch _
func GetAllBranch() (req *[]Branchs) {
	o := orm.NewOrm()
	reqGet := &[]Branchs{}
	o.QueryTable("branchs").RelatedSel().All(reqGet)
	return reqGet
}

//GetDepartByID _
func GetDepartByID(ID int) (dept *Departs, errRet error) {
	reqGet := &Departs{}
	o := orm.NewOrm()
	o.QueryTable("departs").Filter("ID", ID).RelatedSel().One(reqGet)
	return reqGet, errRet
}

//GetDepartList _
func GetDepartList(term, branchID string, limit int) (dept *[]Departs, rowCount int, errRet error) {
	reqGet := &[]Departs{}
	o := orm.NewOrm()
	qs := o.QueryTable("departs")
	cond := orm.NewCondition()
	cond1 := cond.Or("Name__icontains", term)
	if branchID != "" {
		cond1 = cond.Or("Name__icontains", term).And("Branch__ID", branchID)
	}
	qs.SetCond(cond1).RelatedSel().Limit(limit).All(reqGet)
	return reqGet, len(*reqGet), errRet
}

//CreateDeparts _
func CreateDeparts(depart Departs) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&depart)
	o.Commit()
	if err == nil {
		retID = id
	}
	return retID, err
}

//UpdateDeparts _
func UpdateDeparts(depart Departs) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetDepartByID(depart.ID)
	if getUpdate == nil {
		errRet = errors.New("ไม่พบข้อมูล")
	} else {
		depart.CreatedAt = getUpdate.CreatedAt
		if num, errUpdate := o.Update(&depart); errUpdate != nil {
			errRet = errUpdate
			_ = num
		}
	}
	return errRet
}

//DeleteDepartByID _
func DeleteDepartByID(ID int) (errRet error) {
	o := orm.NewOrm()
	if num, errUpdate := o.Delete(&Departs{ID: ID}); errUpdate != nil {
		errRet = errUpdate
		_ = num
	}
	return errRet
}
