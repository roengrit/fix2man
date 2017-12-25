package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//User เก็บข้อมูล User ใช้งาน
type Users struct {
	ID        int
	Username  string
	Password  string
	Email     string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Users))
}
