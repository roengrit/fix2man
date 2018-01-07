package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Units _
type Units struct {
	ID int
	//Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	orm.RegisterModel(
		new(Units),
	) // Need to register model in init
}
