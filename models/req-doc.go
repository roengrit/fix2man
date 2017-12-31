package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//RequestDocument _
type RequestDocument struct {
	ID           int
	DocNo        string      `orm:"size(20)"`
	User         *Users      `orm:"rel(one)"`
	Tel          string      `orm:"size(50)"`
	Branch       *Branchs    `orm:"null;rel(one)"`
	Depart       *Departs    `orm:"null;rel(one)"`
	Building     *Buildings  `orm:"null;rel(one)"`
	Class        *Class      `orm:"null;rel(one)"`
	Room         *Rooms      `orm:"null;rel(one)"`
	Equipmen     *Equipments `orm:"null;rel(one)"`
	Location     string      `orm:"size(225)"`
	SerailNumber string      `orm:"size(50)"`
	Details      string      `orm:"size(500)"`
	Remark       string      `orm:"size(300)"`
	CreatedAt    time.Time   `orm:"now()`
	UpdatedAt    time.Time
}

func init() {
	orm.RegisterModel(new(RequestDocument)) // Need to register model in init
}
