package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//NormalModel _
type NormalModel struct {
	RetOK      bool
	RetCount   int64
	RetData    string
	ID         int64
	Name       string
	XSRF       string
	FlagAction string
	ListData   interface{}
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

//GetListEntity _
func GetListEntity(entity, top, term string) (num int64, err error, entityList []NormalEntity) {
	var sql = "SELECT i_d,code, name,lock FROM " + entity + " WHERE lower(name) like lower(?) or lower(code) like lower(?) order by code limit {0}"
	if top == "0" {
		sql = strings.Replace(sql, "limit {0}", "", -1)
	} else {
		sql = strings.Replace(sql, "{0}", top, -1)
	}
	o := orm.NewOrm()
	num, err = o.Raw(sql, "%"+term+"%", "%"+term+"%").QueryRows(&entityList)
	return num, err, entityList
}

//GetListEntityWithParent _
func GetListEntityWithParent(entity, entityParent, top, parentID, term string) (num int64, err error, entityList []NormalEntity) {
	var sql = "SELECT i_d, name FROM " + entity + " WHERE " + entityParent + "= ? and lower(name) like lower(?) order by name limit {0}"
	if top == "0" {
		sql = strings.Replace(sql, "limit {0}", "", -1)
	} else {
		sql = strings.Replace(sql, "{0}", top, -1)
	}
	o := orm.NewOrm()
	num, err = o.Raw(sql, parentID, "%"+term+"%").QueryRows(&entityList)
	return num, err, entityList
}

//GetEntity _
func GetEntity(entity, id string) (err error, entityItem NormalEntity) {
	var sql = "SELECT i_d,code, name,lock FROM " + entity + " WHERE i_d = ?  limit 1"
	o := orm.NewOrm()
	err = o.Raw(sql, id).QueryRow(&entityItem)
	return err, entityItem
}

//GetEntityByCode _
func GetEntityByCode(entity, code string) (err error, entityItem NormalEntity) {
	var sql = "SELECT i_d,code, name,lock FROM " + entity + " WHERE code = ?  limit 1"
	o := orm.NewOrm()
	err = o.Raw(sql, code).QueryRow(&entityItem)
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
func CreateEntity(entity, code, name string) (retID int64, errRet error) {
	err, ret := GetEntityByCode(entity, code)
	retID = 0
	if err == nil && (NormalEntity{}) != ret {
		err = errors.New("รหัสซ้ำ กรุณากำหนดรหัสใหม่")
	} else {
		var sql = "INSERT INTO " + entity + " (code,name,created_at,updated_at) values(?,?,now(),now()) "
		o := orm.NewOrm()
		res, errInsert := o.Raw(sql, code, name).Exec()
		err = errInsert
		if errInsert != nil {
			retID, _ = res.LastInsertId()
		}
	}
	return retID, err
}

//UpdateEntity _
func UpdateEntity(entity, id, code, name string) (err error) {
	err, ret := GetEntity(entity, id)
	if err == nil && (NormalEntity{}) != ret { //พบข้อมูลที่ต้องการแก้ไข
		if ret.Lock { //ข้อมูล ล็อค  ไม่อนุญาตให้แก้ไข หรือ ลบข้อมูล
			err = errors.New("ไม่อนุญาตให้แก้ไข หรือ ลบข้อมูล")
		} else {
			err, ret = GetEntityByCode(entity, code)
			if err == nil && (NormalEntity{}) != ret { //พบข้อมูลมีรหัสซ้ำ
				if num, _ := strconv.ParseInt(id, 10, 64); ret.ID != int(num) { //ข้อมูลมีรหัสซ้ำ ไม่ใช่ข้อมูลตัวเอง
					err = errors.New("รหัสซ้ำ กรุณากำหนดรหัสใหม่")
				}
			} else {
				if err == orm.ErrNoRows {
					err = nil
				}
			}
		}
	} else {
		err = errors.New("ไม่พบข้อมูล")
	}
	fmt.Println(err)
	if err == nil {
		var sql = "UPDATE " + entity + " SET code = ? , name = ? WHERE i_d = ?"
		o := orm.NewOrm()
		_, err = o.Raw(sql, code, name, id).Exec()
	}
	return err
}

//GetMaxEntity _
func GetMaxEntity(entity string) (code string) {

	var lists []orm.ParamsList
	var sql = "SELECT COALESCE(MAX(code),'0000') AS code FROM " + entity + " WHERE code NOT LIKE '%[^0-9]%' AND code != '' AND LENGTH(code) = 4"
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
