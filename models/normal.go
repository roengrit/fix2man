package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//NormalModel _
type NormalModel struct {
	RetOK    bool
	RetCount int64
	RetData  string
}

// NormalEntity _
type NormalEntity struct {
	ID        int
	Code      string
	Name      string
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Roles _
type Roles struct {
	ID        int
	Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Status _
type Status struct {
	ID        int
	Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Units _
type Units struct {
	ID        int
	Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Branchs _
type Branchs struct {
	ID        int
	Code      string `orm:"size(20)"`
	Name      string `orm:"size(225)"`
	Lock      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CodeOnly struct {
	MaxCode string
}

func init() {
	orm.RegisterModel(new(Roles), new(Status), new(Units), new(Branchs)) // Need to register model in init
}

//GetListEntity _
func GetListEntity(entity, top, term string) (num int64, err error, entityList []NormalEntity) {
	var sql = "SELECT i_d,code, name,lock FROM " + entity + " WHERE name like ? or code like ? order by code limit {0}"
	if top == "0" {
		sql = strings.Replace(sql, "limit {0}", "", -1)
	} else {
		sql = strings.Replace(sql, "{0}", top, -1)
	}
	o := orm.NewOrm()
	num, err = o.Raw(sql, "%"+term+"%", "%"+term+"%").QueryRows(&entityList)
	return num, err, entityList
}

//GetMaxEntity _
func GetMaxEntity(entity string) (code string) {

	var lists []orm.ParamsList
	var sql = "SELECT COALESCE(MAX(code),'0001') AS code FROM " + entity + " WHERE code NOT LIKE '%[^0-9]%' AND code != '' AND LENGTH(code) = 4"
	o := orm.NewOrm()
	num, err := o.Raw(sql).ValuesList(&lists)
	if err == nil && num > 0 {
		max := lists[0]
		maxVal, _ := strconv.ParseInt(max[0].(string), 10, 64)
		code = fmt.Sprintf("%04d", maxVal+1)
	} else {
		code = "0001"
	}
	return code
}
