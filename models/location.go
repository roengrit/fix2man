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
	CreatedAt time.Time `orm:"auto_now_add"`
	UpdatedAt time.Time `orm:"null"`
}

//Rooms _
type Rooms struct {
	ID        int
	Class     *Class `orm:"rel(one)"`
	Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time `orm:"auto_now_add"`
	UpdatedAt time.Time `orm:"null"`
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

//GetBranchByID _
func GetBranchByID(ID int) (branch *Branchs, errRet error) {
	reqGet := &Branchs{}
	o := orm.NewOrm()
	o.QueryTable("branchs").Filter("ID", ID).RelatedSel().One(reqGet)
	return reqGet, errRet
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

//GetBuildingsByID _
func GetBuildingsByID(ID int) (dept *Buildings, errRet error) {
	reqGet := &Buildings{}
	o := orm.NewOrm()
	o.QueryTable("buildings").Filter("ID", ID).RelatedSel().One(reqGet)
	return reqGet, errRet
}

//GetBuildingsList _
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

//CreateBuildings _
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

//UpdateBuildings _
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

//DeleteBuildingsByID _
func DeleteBuildingsByID(ID int) (errRet error) {
	o := orm.NewOrm()
	if num, errUpdate := o.Delete(&Buildings{ID: ID}); errUpdate != nil {
		errRet = errUpdate
		_ = num
	}
	return errRet
}

//////////////////////////////////////////////// ชั้น //////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

//GetClassByID _
func GetClassByID(ID int) (dept *Class, errRet error) {
	reqGet := &Class{}
	o := orm.NewOrm()
	o.QueryTable("class").Filter("ID", ID).RelatedSel().One(reqGet)
	return reqGet, errRet
}

//GetClassList _
func GetClassList(term, branchID, BuildingID string, limit int) (class *[]Class, rowCount int, errRet error) {
	reqGet := &[]Class{}
	o := orm.NewOrm()
	qs := o.QueryTable("class")
	cond := orm.NewCondition()
	cond1 := cond.Or("Name__icontains", term)
	if branchID != "" {
		cond1 = cond1.And("Building__Branch__ID", branchID)
	}
	if BuildingID != "" {
		cond1 = cond1.And("Building__ID", BuildingID)
	}
	qs.SetCond(cond1).RelatedSel().Limit(limit).All(reqGet)
	return reqGet, len(*reqGet), errRet
}

//CreateClass _
func CreateClass(class Class) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&class)
	o.Commit()
	if err == nil {
		retID = id
	}
	return retID, err
}

//UpdateClass _
func UpdateClass(class Class) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetDepartByID(class.ID)
	if getUpdate == nil {
		errRet = errors.New("ไม่พบข้อมูล")
	} else {
		class.CreatedAt = getUpdate.CreatedAt
		if num, errUpdate := o.Update(&class); errUpdate != nil {
			errRet = errUpdate
			_ = num
		}
	}
	return errRet
}

//DeleteClassByID _
func DeleteClassByID(ID int) (errRet error) {
	o := orm.NewOrm()
	if num, errUpdate := o.Delete(&Class{ID: ID}); errUpdate != nil {
		errRet = errUpdate
		_ = num
	}
	return errRet
}

//////////////////////////////////////////////// ห้อง //////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

//GetRoomByID _
func GetRoomByID(ID int) (dept *Rooms, errRet error) {
	reqGet := &Rooms{}
	o := orm.NewOrm()
	o.QueryTable("rooms").Filter("ID", ID).RelatedSel().One(reqGet)
	return reqGet, errRet
}

//GetRoomList _
func GetRoomList(term, branchID, BuildingID, ClassID string, limit int) (room *[]Rooms, rowCount int, errRet error) {
	reqGet := &[]Rooms{}
	o := orm.NewOrm()
	qs := o.QueryTable("rooms")
	cond := orm.NewCondition()
	cond1 := cond.Or("Name__icontains", term)
	if branchID != "" {
		cond1 = cond1.And("Class__Building__Branch__ID", branchID)
	}
	if BuildingID != "" {
		cond1 = cond1.And("Class__Building__ID", BuildingID)
	}
	if ClassID != "" {
		cond1 = cond1.And("Class__ID", ClassID)
	}
	qs.SetCond(cond1).RelatedSel().Limit(limit).All(reqGet)
	return reqGet, len(*reqGet), errRet
}

//CreateRoom _
func CreateRoom(room Rooms) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&room)
	o.Commit()
	if err == nil {
		retID = id
	}
	return retID, err
}

//UpdateRoom _
func UpdateRoom(room Rooms) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetRoomByID(room.ID)
	if getUpdate == nil {
		errRet = errors.New("ไม่พบข้อมูล")
	} else {
		room.CreatedAt = getUpdate.CreatedAt
		if num, errUpdate := o.Update(&room); errUpdate != nil {
			errRet = errUpdate
			_ = num
		}
	}
	return errRet
}

//DeleteRoomByID _
func DeleteRoomByID(ID int) (errRet error) {
	o := orm.NewOrm()
	if num, errUpdate := o.Delete(&Rooms{ID: ID}); errUpdate != nil {
		errRet = errUpdate
		_ = num
	}
	return errRet
}
