package models

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Provinces _
type Provinces struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NormalModel _
type NormalModel struct {
	RetOK      bool
	RetCount   int64
	RetData    string
	ID         int64
	Name       string
	Del        string
	Title      string
	Alert      string
	XSRF       string
	FlagAction string
	ListData   interface{}
	Data1      interface{}
	Data2      interface{}
	Data3      interface{}
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

func init() {
	orm.RegisterModel(new(Provinces)) // Need to register model in init
}

//GetAllProvince _
func GetAllProvince() (req *[]Provinces) {
	o := orm.NewOrm()
	reqGet := &[]Provinces{}
	o.QueryTable("provinces").RelatedSel().OrderBy("ID").All(reqGet)
	return reqGet
}

//GetListEntity _
func GetListEntity(entity string, top int, term string) (num int64, err error, entityList []NormalEntity) {
	var sql = "SELECT i_d,name,lock FROM " + entity + " WHERE lower(name) like lower(?) order by i_d limit {0}"
	if top == 0 {
		sql = strings.Replace(sql, "limit {0}", "", -1)
	} else {
		sql = strings.Replace(sql, "{0}", strconv.Itoa(top), -1)
	}
	o := orm.NewOrm()
	num, err = o.Raw(sql, "%"+term+"%").QueryRows(&entityList)
	return num, err, entityList
}

//GetListEntityWithParent _
func GetListEntityWithParent(entity string, entityParent string, top int, parentID string, term string) (num int64, err error, entityList []NormalEntity) {
	var sql = "SELECT i_d, name FROM " + entity + " WHERE " + entityParent + "= ? and lower(name) like lower(?) order by name limit {0}"
	if top == 0 {
		sql = strings.Replace(sql, "limit {0}", "", -1)
	} else {
		sql = strings.Replace(sql, "{0}", strconv.Itoa(top), -1)
	}
	o := orm.NewOrm()
	num, err = o.Raw(sql, parentID, "%"+term+"%").QueryRows(&entityList)
	return num, err, entityList
}

//GetEntity _
func GetEntity(entity, id string) (err error, entityItem NormalEntity) {
	var sql = "SELECT i_d,name,lock FROM " + entity + " WHERE i_d = ?  limit 1"
	o := orm.NewOrm()
	err = o.Raw(sql, id).QueryRow(&entityItem)
	return err, entityItem
}

//DeleteEntity _
func DeleteEntity(entity, id string) (errRet error) {
	err, ret := GetEntity(entity, id)
	if err == nil && (NormalEntity{}) != ret {
		if ret.Lock {
			err = errors.New("ไม่อนุญาตให้แก้ไข หรือ ลบข้อมูล")
		}
	} else {
		err = errors.New("ไม่พบข้อมูล")
	}
	if err == nil {
		var sql = "DELETE FROM " + entity + " WHERE i_d = ?  "
		o := orm.NewOrm()
		_, err = o.Raw(sql, id).Exec()
	}
	errRet = err
	return errRet
}

//CreateEntity _
func CreateEntity(entity, name string) (retID int64, errRet error) {
	retID = 0
	var sql = "INSERT INTO " + entity + " (name,created_at,updated_at) values(?,now(),now()) "
	o := orm.NewOrm()
	res, errInsert := o.Raw(sql, name).Exec()
	errRet = errInsert
	if errInsert != nil {
		retID, _ = res.LastInsertId()
	}
	return retID, errRet
}

//UpdateEntity _
func UpdateEntity(entity, id, name string) (err error) {
	err, ret := GetEntity(entity, id)
	if err == nil && (NormalEntity{}) != ret { //พบข้อมูลที่ต้องการแก้ไข
		if ret.Lock { //ข้อมูล ล็อค  ไม่อนุญาตให้แก้ไข หรือ ลบข้อมูล
			err = errors.New("ไม่อนุญาตให้แก้ไข หรือ ลบข้อมูล")
		}
	} else {
		err = errors.New("ไม่พบข้อมูล")
	}
	if err == nil {
		var sql = "UPDATE " + entity + " SET name = ? WHERE i_d = ?"
		o := orm.NewOrm()
		_, err = o.Raw(sql, name, id).Exec()
	}
	return err
}

//GetMaxDoc
func GetMaxDoc(entity, format string) (docno string) {

	var lists []orm.ParamsList
	var sql = `SELECT 
				 concat( '` + format + `' , 
					  date_part( 'year', CURRENT_DATE :: timestamp ) :: text ,
					  case WHEN length( date_part( 'month', CURRENT_DATE :: timestamp ) :: text )  = 1 then  
								concat('0', date_part( 'month', CURRENT_DATE :: timestamp ) :: text ) else 
										  date_part( 'month', CURRENT_DATE :: timestamp ) :: text 
					  end  , 
					  '-',
					  LPAD((COALESCE( MAX ( SUBSTRING ( doc_no FROM '[0-9]{5,}$' )), '0' ) :: NUMERIC + 1)  :: text, 5, '0')
					  ) AS  doc_no 
				 FROM
					 ` + entity + `
				 WHERE
					 doc_no LIKE'` + format + `%' 	 
					 AND doc_date BETWEEN date_trunc( 'month', CURRENT_DATE ) :: DATE  AND ( date_trunc( 'month', CURRENT_DATE ) + INTERVAL '1 month - 1 day' ) :: DATE`
	o := orm.NewOrm()
	num, err := o.Raw(sql).ValuesList(&lists)
	if err == nil && num > 0 {
		max := lists[0]
		maxVal := max[0].(string)
		docno = maxVal
	} else {
		docno = ""
	}
	return docno
}
