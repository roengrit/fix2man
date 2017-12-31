package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Branchs _
type Branchs struct {
	ID                  int
	Code                string `orm:"size(20)"`
	Name                string `orm:"size(225)"`
	Lock                bool
	TokenLine           string
	CostAvgPerTechnical float64
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func init() {
	orm.RegisterModel(
		new(Branchs),
	) // Need to register model in init
}
