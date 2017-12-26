package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
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

// Login getUser Pass
func Login(username, password string) (ok bool, errRet string) {
	o := orm.NewOrm()
	user := Users{Username: username}
	err := o.Read(&user, "Username")
	if err == orm.ErrNoRows {
		errRet = "รหัสผู้ใช้/รหัสผ่านไม่ถูกต้อง"
	} else {
		if errCript := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); errCript != nil {
			errRet = "รหัสผู้ใช้/รหัสผ่านไม่ถูกต้อง"
		} else {
			ok = true
		}
	}
	return ok, errRet
}
