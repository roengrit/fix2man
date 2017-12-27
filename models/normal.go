package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//NormalModel _
type NormalModel struct {
	RetOK    bool
	RetCount int64
	RetData  string
}

type NormalEntity struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Roles struct {
	ID        int
	Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Status struct {
	ID        int
	Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Units struct {
	ID        int
	Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Branchs struct {
	ID        int
	Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	orm.RegisterModel(new(Roles), new(Status), new(Units), new(Branchs)) // Need to register model in init
}
