package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Departs _
type Departs struct {
	ID     int
	Branch *Branchs `orm:"rel(one)"`
	//Code      string   `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	orm.RegisterModel(
		new(Departs),
	)
}
