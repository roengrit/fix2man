package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

//Suppliers _
type Suppliers struct {
	ID        int
	Name      string     `orm:"size(225)"`
	Address   string     `orm:"size(300)"`
	Province  *Provinces `orm:"rel(one)"`
	PostCode  string     `orm:"size(10)"`
	Contact   string     `orm:"size(255)"`
	Tel       string     `orm:"size(100)"`
	Remark    string     `orm:"size(100)"`
	CreatedAt time.Time  `orm:"auto_now_add"`
	UpdatedAt time.Time  `orm:"null"`
}

func init() {
	orm.RegisterModel(new(Suppliers))
}

//CreateSuppliers _
func CreateSuppliers(sup Suppliers) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&sup)
	o.Commit()
	if err == nil {
		retID = id
	}
	return retID, err
}

//UpdateSuppliers _
func UpdateSuppliers(sup Suppliers) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetSuppliers(sup.ID)
	if getUpdate == nil {
		errRet = errors.New("ไม่พบข้อมูล")
	} else {
		sup.CreatedAt = getUpdate.CreatedAt
		if num, errUpdate := o.Update(&sup); errUpdate != nil {
			errRet = errUpdate
			_ = num
		}
	}
	return errRet
}

//GetSuppliers _
func GetSuppliers(ID int) (sup *Suppliers, errRet error) {
	reqGet := &Suppliers{}
	o := orm.NewOrm()
	o.QueryTable("suppliers").Filter("ID", ID).RelatedSel().One(reqGet)
	return reqGet, errRet
}

//GetSuppliersList _
func GetSuppliersList(term string, limit int) (sup *[]Suppliers, errRet error) {
	orm.Debug = true
	reqGet := &[]Suppliers{}
	o := orm.NewOrm()
	qs := o.QueryTable("suppliers")
	cond := orm.NewCondition()
	cond1 := cond.Or("Name__icontains", term).
		Or("Tel__icontains", term).
		Or("Contact__icontains", term).
		Or("Remark__icontains", term).
		Or("Address__icontains", term)
	qs.SetCond(cond1).RelatedSel().Limit(limit).All(reqGet)
	return reqGet, errRet
}
