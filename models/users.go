package models

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

//Users เก็บข้อมูล User ใช้งาน
type Users struct {
	ID               int
	Username         string `orm:"size(50)"`
	Password         string `orm:"size(255)"`
	Name             string `orm:"size(500)"`
	Tel              string `orm:"size(255)"`
	CostPerTechnical float64
	Roles            *Roles     `orm:"rel(fk)"`
	Branch           *Branchs   `orm:"null;rel(one)"`
	Depart           *Departs   `orm:"null;rel(one)"`
	Building         *Buildings `orm:"null;rel(one)"`
	Rooms            *Rooms     `orm:"null;rel(one)"`
	Class            *Class     `orm:"null;rel(one)"`
	Active           bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
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

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

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

//GetUser _
func GetUser(username string) (ok bool, errRet string) {
	o := orm.NewOrm()
	user := Users{Username: username}
	err := o.Read(&user, "Username")
	if err == orm.ErrNoRows {
		errRet = "ไม่พบ email นี้ในระบบ"
	} else {
		ok = true
	}
	return ok, errRet
}

//GetUserByID _
func GetUserByID(ID string) (user *Users, errRet string) {
	o := orm.NewOrm()
	id, _ := strconv.Atoi(ID)
	//user = Users{ID: id}
	userGet := &Users{}
	o.QueryTable("users").Filter("ID", id).RelatedSel().One(userGet)
	if nil != userGet {
		userGet.Password = ""
		userGet.Roles = nil
	}
	return userGet, errRet
}

//GetUserByUserName _
func GetUserByUserName(username string) (user *Users, errRet string) {
	userGet := &Users{}
	o := orm.NewOrm()
	o.QueryTable("users").Filter("Username", username).RelatedSel().One(userGet)
	if nil != userGet {
		userGet.Password = ""
		userGet.Roles = nil
	}
	return userGet, errRet
}

//GetUserList _
func GetUserList(top, term string) (num int64, err error, userList []Users) {
	var sql = "SELECT i_d,name FROM users WHERE lower(name) like lower(?) order by name limit {0}"
	if top == "0" {
		sql = strings.Replace(sql, "limit {0}", "", -1)
	} else {
		sql = strings.Replace(sql, "{0}", top, -1)
	}
	o := orm.NewOrm()
	num, err = o.Raw(sql, "%"+term+"%").QueryRows(&userList)
	return num, err, userList
}

//ForgetPass _
func ForgetPass(username, newPass string) (ok bool, errRet string) {
	o := orm.NewOrm()
	user := Users{Username: username}
	err := o.Read(&user, "Username")
	if err == orm.ErrNoRows {
		errRet = "ไม่พบ email นี้ในระบบ"
	} else {

		if hash, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost); err != nil {
			errRet = err.Error()
		} else {
			user.Password = string(hash)
			if num, errUpdate := o.Update(&user); errUpdate != nil {
				errRet = errUpdate.Error()
				_ = num
			} else {
				ok = true
			}
		}
	}
	return ok, errRet
}

//RandStringRunes password _
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
