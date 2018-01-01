package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Status _
type Status struct {
	ID        int
	Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	IsDef     bool
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	orm.RegisterModel(
		new(Status),
	) // Need to register model in init
}

//GetAllStatus _
func GetAllStatus() (req *[]Status) {
	o := orm.NewOrm()
	reqGet := &[]Status{}
	o.QueryTable("status").RelatedSel().All(reqGet)
	return reqGet
}
