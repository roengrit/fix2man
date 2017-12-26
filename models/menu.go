package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Menu _
type Menu struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	orm.RegisterModel(new(Menu)) //Need to register model in init
}
