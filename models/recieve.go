package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

//Receive _
type Receive struct {
	ID             int
	Flag           int
	Active         bool
	DocType        int
	DocNo          string    `orm:"size(30)"`
	DocDate        time.Time `form:"-" orm:"null"`
	DocTime        string    `orm:"size(6)"`
	DocRefNo       string    `orm:"size(30)"`
	TableNo        string    `orm:"size(300)"`
	Member         *Member   `orm:"rel(fk)"`
	MemberName     string    `orm:"size(300)"`
	DiscountType   int
	DiscountWord   string  `orm:"size(300)"`
	TotalDiscount  float64 `orm:"digits(12);decimals(2)"`
	TotalAmount    float64 `orm:"digits(12);decimals(2)"`
	TotalNetAmount float64 `orm:"digits(12);decimals(2)"`
	CreditDay      int
	CreditDate     time.Time    `orm:"type(date)"`
	Remark         string       `orm:"size(300)"`
	CancelRemark   string       `orm:"size(300)"`
	Creator        *Users       `orm:"rel(fk)"`
	CreatedAt      time.Time    `orm:"auto_now_add;type(datetime)"`
	Editor         *Users       `orm:"null;rel(fk)"`
	EditedAt       time.Time    `orm:"null;auto_now;type(datetime)"`
	CancelUser     *Users       `orm:"null;rel(fk)"`
	CancelAt       time.Time    `orm:"null;type(datetime)"`
	ReceiveSub     []ReceiveSub `orm:"-"`
}

//ReceiveSub _
type ReceiveSub struct {
	ID          int
	Flag        int
	Active      bool
	DocNo       string    `orm:"size(30)"`
	DocDate     time.Time `form:"-" orm:"null"`
	Product     *Product  `orm:"rel(fk)"`
	Serial      string    `orm:"size(100)"`
	Unit        *Unit     `orm:"rel(fk)"`
	Qty         float64   `orm:"digits(12);decimals(2)"`
	RemainQty   float64   `orm:"digits(12);decimals(2)"`
	AverageCost float64   `orm:"digits(12);decimals(2)"`
	Price       float64   `orm:"digits(12);decimals(2)"`
	TotalPrice  float64   `orm:"digits(12);decimals(2)"`
	Creator     *Users    `orm:"rel(fk)"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Receive), new(ReceiveSub)) // Need to register model in init
}

//CreateReceive _
func CreateReceive(receive Receive, user Users) (retID int64, errRet error) {
	receive.DocNo = GetMaxDoc("receive", "REC")
	receive.Creator = &user
	receive.CreatedAt = time.Now()
	receive.CreditDay = 0
	receive.CreditDate = time.Now()
	receive.Active = true
	var fullDataSub []ReceiveSub
	for _, val := range receive.ReceiveSub {
		if val.Product.ID != 0 {
			val.CreatedAt = time.Now()
			val.Creator = &user
			val.DocNo = receive.DocNo
			val.Flag = receive.Flag
			val.Active = true
			val.DocDate = receive.DocDate
			fullDataSub = append(fullDataSub, val)
		}
	}
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&receive)
	id, err = o.InsertMulti(len(fullDataSub), fullDataSub)
	if err == nil {
		retID = id
		o.Commit()
	} else {
		o.Rollback()
	}
	errRet = err
	return retID, errRet
}

//UpdateReceive _
func UpdateReceive(receive Receive, user Users) (retID int64, errRet error) {
	docCheck, _ := GetReceive(receive.ID)
	if docCheck.ID == 0 {
		errRet = errors.New("ไม่พบข้อมูล")
	}
	receive.Creator = docCheck.Creator
	receive.CreatedAt = docCheck.CreatedAt
	receive.CreditDay = docCheck.CreditDay
	receive.CreditDate = docCheck.CreditDate
	receive.EditedAt = time.Now()
	receive.Editor = &user
	receive.Active = docCheck.Active
	var fullDataSub []ReceiveSub
	for _, val := range receive.ReceiveSub {
		if val.Product.ID != 0 {
			val.CreatedAt = time.Now()
			val.Creator = &user
			val.DocNo = receive.DocNo
			val.Flag = receive.Flag
			val.Active = receive.Active
			val.DocDate = receive.DocDate
			fullDataSub = append(fullDataSub, val)
		}
	}
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Update(&receive)
	if err == nil {
		_, err = o.QueryTable("receive_sub").Filter("doc_no", receive.DocNo).Delete()
	}
	if err == nil {
		_, err = o.InsertMulti(len(fullDataSub), fullDataSub)
	}
	if err == nil {
		retID = id
		o.Commit()
	} else {
		o.Rollback()
	}
	errRet = err
	return retID, errRet
}

//GetReceive _
func GetReceive(ID int) (doc *Receive, errRet error) {
	receiveDoc := &Receive{}
	o := orm.NewOrm()
	o.QueryTable("receive").Filter("ID", ID).RelatedSel().One(receiveDoc)
	o.QueryTable("receive_sub").Filter("doc_no", receiveDoc.DocNo).RelatedSel().All(&receiveDoc.ReceiveSub)
	doc = receiveDoc
	return doc, errRet
}

//GetReceiveList _
func GetReceiveList(term string, limit int, dateBegin, dateEnd string) (sup *[]Receive, rowCount int, errRet error) {
	receive := &[]Receive{}
	o := orm.NewOrm()
	qs := o.QueryTable("receive")
	condSub1 := orm.NewCondition()
	condSub2 := orm.NewCondition()
	cond1 := condSub1.And("doc_date__gte", dateBegin).And("doc_date__lte", dateEnd)
	qs = qs.SetCond(cond1)
	if dateBegin != "" && dateEnd != "" {
		cond2 := condSub2.Or("Member__Name__icontains", term).Or("DocNo__icontains", term).Or("Remark__icontains", term)
		cond1 = cond1.AndCond(cond2)
		qs = qs.SetCond(cond1)
	}
	qs.RelatedSel().Limit(limit).All(receive)
	return receive, len(*receive), errRet
}

//UpdateCancelReceive _
func UpdateCancelReceive(ID int, remark string, user Users) (retID int64, errRet error) {
	docCheck := &Receive{}
	o := orm.NewOrm()
	o.QueryTable("receive").Filter("ID", ID).RelatedSel().One(docCheck)
	if docCheck.ID == 0 {
		errRet = errors.New("ไม่พบข้อมูล")
	}
	docCheck.Active = false
	docCheck.CancelRemark = remark
	docCheck.CancelAt = time.Now()
	docCheck.CancelUser = &user
	o.Begin()
	_, err := o.Update(docCheck)
	if err == nil {
		_, err = o.Raw("update receive_sub set active = false where doc_no = ?", docCheck.DocNo).Exec()
	}
	if err != nil {
		o.Rollback()
	} else {
		o.Commit()
	}
	errRet = err
	return retID, errRet
}
