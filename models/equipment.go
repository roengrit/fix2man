package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Equipments _
type Equipments struct {
	ID        int
	Name      string
	CreatedAt time.Time `orm:"now()`
	UpdatedAt time.Time
}

func init() {
	orm.RegisterModel(new(Equipments)) // Need to register model in init
}
