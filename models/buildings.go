package models

import (
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Buildings _
type Buildings struct {
	ID     int
	Branch *Branchs `orm:"rel(fk)"`
	//Code      string   `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Class _
type Class struct {
	ID       int
	Building *Buildings `orm:"rel(one)"`
	//Code      string     `orm:"size(20)"`
	Name      string `orm:"size(225)"`
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
	) // Need to register model in init
}

func GetBuildingList(top, branchID, term string) (num int64, err error, buildList []Buildings) {
	var sql = "SELECT i_d,name FROM buildings WHERE branch_id = ? and lower(name) like lower(?) order by name limit {0}"
	if top == "0" {
		sql = strings.Replace(sql, "limit {0}", "", -1)
	} else {
		sql = strings.Replace(sql, "{0}", top, -1)
	}
	o := orm.NewOrm()
	num, err = o.Raw(sql, branchID, "%"+term+"%").QueryRows(&buildList)
	return num, err, buildList
}
