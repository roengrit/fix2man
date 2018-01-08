package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Products _
type Products struct {
	ID        int
	Name      string
	CreatedAt time.Time `orm:"now()`
	UpdatedAt time.Time
}

//Units _
type Units struct {
	ID        int
	Name      string `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	orm.RegisterModel(new(Products), new(Units)) // Need to register model in init
}
