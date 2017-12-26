package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Users เก็บข้อมูล User ใช้งาน
type Users struct {
	ID        int
	Username  string
	Password  string
	Email     string
	RoleID    int
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Role เก็บข้อมูลสิทธิ์ใช้งาน
type Role struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Permiss เก็บข้อมูลสิทธิ์ใช้งาน
type Permiss struct {
	ID        int
	RoleID    int
	MenuID    int
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	orm.RegisterModel(new(Users), new(Permiss)) // Need to register model in init
}
