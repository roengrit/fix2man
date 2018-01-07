package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Status _
type Status struct {
	ID int
	//Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	IsDef     bool
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	orm.RegisterModel(
		new(Status),
	)
}

//GetAllStatus _
func GetAllStatus() (req *[]Status) {
	o := orm.NewOrm()
	reqGet := &[]Status{}
	o.QueryTable("status").RelatedSel().OrderBy("ID").All(reqGet)
	return reqGet
}

//GetAllStatusExcludeId _
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
