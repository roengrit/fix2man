package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Buildings _
type Buildings struct {
	ID        int
	Branch    *Branchs `orm:"rel(fk)"`
	Code      string   `orm:"size(20)"`
	Name      string   `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Class _
type Class struct {
	ID        int
	Building  *Buildings `orm:"rel(fk)"`
	Code      string     `orm:"size(20)"`
	Name      string     `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Rooms _
type Rooms struct {
	ID        int
	Class     *Class `orm:"rel(fk)"`
	Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Locations _
type Locations struct {
	ID        int
	Room      *Rooms `orm:"rel(fk)"`
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
		new(Locations),
	) // Need to register model in init
}
