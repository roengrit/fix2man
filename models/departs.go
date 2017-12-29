package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Departs _
type Departs struct {
	ID        int
	Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	Lock      bool
	DepartMap []*DepartMaps `orm:"reverse(many)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//DepartMaps _
type DepartMaps struct {
	ID        int
	Branch    *Branchs   `orm:"rel(fk)"`
	Depart    *Departs   `orm:"rel(fk)"`
	Building  *Buildings `orm:"rel(fk)"`
	Class     *Class     `orm:"rel(fk)"`
	Room      *Rooms     `orm:"rel(fk)"`
	Code      string     `orm:"size(20)"`
	Name      string     `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	orm.RegisterModel(
		new(Departs),
		new(DepartMaps),
	) // Need to register model in init
}
