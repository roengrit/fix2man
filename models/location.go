package models

import (
	"errors"
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
	Branch    *Branchs `orm:"rel(one)"`
	Name      string   `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time `orm:"auto_now_add"`
	UpdatedAt time.Time `orm:"null"`
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

//////////////////////////////////////////////// อาคาร //////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

//GetDepartByID _
func GetBuildingsByID(ID int) (dept *Buildings, errRet error) {
	reqGet := &Buildings{}
	o := orm.NewOrm()
	o.QueryTable("buildings").Filter("ID", ID).RelatedSel().One(reqGet)
	return reqGet, errRet
}

//GetDepartList _
func GetBuildingsList(term, branchID string, limit int) (buildings *[]Buildings, rowCount int, errRet error) {
	reqGet := &[]Buildings{}
	o := orm.NewOrm()
	qs := o.QueryTable("buildings")
	cond := orm.NewCondition()
	cond1 := cond.Or("Name__icontains", term)
	if branchID != "" {
		cond1 = cond.Or("Name__icontains", term).And("Branch__ID", branchID)
	}
	qs.SetCond(cond1).RelatedSel().Limit(limit).All(reqGet)
	return reqGet, len(*reqGet), errRet
}

//CreateDeparts _
func CreateBuildings(buildings Buildings) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&buildings)
	o.Commit()
	if err == nil {
		retID = id
	}
	return retID, err
}

//UpdateDeparts _
func UpdateBuildings(buildings Buildings) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetDepartByID(buildings.ID)
	if getUpdate == nil {
		errRet = errors.New("ไม่พบข้อมูล")
	} else {
		buildings.CreatedAt = getUpdate.CreatedAt
		if num, errUpdate := o.Update(&buildings); errUpdate != nil {
			errRet = errUpdate
			_ = num
		}
	}
	return errRet
}

//DeleteDepartByID _
func DeleteBuildingsByID(ID int) (errRet error) {
	o := orm.NewOrm()
	if num, errUpdate := o.Delete(&Buildings{ID: ID}); errUpdate != nil {
		errRet = errUpdate
		_ = num
	}
	return errRet
}
<<<<<<< HEAD

//////////////////////////////////////////////// ชั้น //////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

//GetDepartByID _
func GetClassByID(ID int) (dept *Buildings, errRet error) {
	reqGet := &Buildings{}
	o := orm.NewOrm()
	o.QueryTable("buildings").Filter("ID", ID).RelatedSel().One(reqGet)
	return reqGet, errRet
}

//GetClassList _
func GetClassList(term, branchID, BuildingID string, limit int) (buildings *[]Buildings, rowCount int, errRet error) {
	reqGet := &[]Buildings{}
	o := orm.NewOrm()
	qs := o.QueryTable("buildings")
	cond := orm.NewCondition()
	cond1 := cond.Or("Name__icontains", term)
	if branchID != "" {
		cond1 = cond1.And("Branch__ID", branchID)
	}
	if branchID != "" {
		cond1 = cond1.And("Building__ID", BuildingID)
	}
	qs.SetCond(cond1).RelatedSel().Limit(limit).All(reqGet)
	return reqGet, len(*reqGet), errRet
}
=======
>>>>>>> 767d01def22501dd3124d96a63527120ed2464ba
