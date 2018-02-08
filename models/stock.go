package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

//StockCount _
type StockCount struct {
	ID             int
	Flag           int
	Active         bool
	FlagTemp       int
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
	CreditDate     time.Time       `orm:"type(date)"`
	Remark         string          `orm:"size(300)"`
	CancelRemark   string          `orm:"size(300)"`
	Creator        *Users          `orm:"rel(fk)"`
	CreatedAt      time.Time       `orm:"auto_now_add;type(datetime)"`
	Editor         *Users          `orm:"null;rel(fk)"`
	EditedAt       time.Time       `orm:"null;auto_now;type(datetime)"`
	CancelUser     *Users          `orm:"null;rel(fk)"`
	CancelAt       time.Time       `orm:"null;type(datetime)"`
	StockCountSub  []StockCountSub `orm:"-"`
}

//StockCountSub _
type StockCountSub struct {
	ID          int
	Flag        int
	Active      bool
	DocNo       string    `orm:"size(30)"`
	DocDate     time.Time `form:"-" orm:"null"`
	Product     *Product  `orm:"rel(fk)"`
	Unit        *Unit     `orm:"rel(fk)"`
	BalanceQty  float64   `orm:"digits(12);decimals(2)"`
	Qty         float64   `orm:"digits(12);decimals(2)"`
	DiffQty     float64   `orm:"digits(12);decimals(2)"`
	RemainQty   float64   `orm:"digits(12);decimals(2)"`
	AverageCost float64   `orm:"digits(12);decimals(2)"`
	Price       float64   `orm:"digits(12);decimals(2)"`
	TotalPrice  float64   `orm:"digits(12);decimals(2)"`
	Remark      string    `orm:"size(300)"`
	Creator     *Users    `orm:"rel(fk)"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(StockCount), new(StockCountSub)) // Need to register model in init
}

//CreateStockCount _
func CreateStockCount(StockCount StockCount, user Users) (retID int64, errRet error) {
	StockCount.DocNo = GetMaxDoc("stock_count", "STK")
	StockCount.Creator = &user
	StockCount.CreatedAt = time.Now()
	StockCount.CreditDay = 0
	StockCount.CreditDate = time.Now()
	StockCount.Active = true
	var fullDataSub []StockCountSub
	for _, val := range StockCount.StockCountSub {
		if val.Product.ID != 0 {
			Product, _ := GetProduct(val.Product.ID)
			val.CreatedAt = time.Now()
			val.Creator = &user
			val.DocNo = StockCount.DocNo
			val.Flag = StockCount.Flag
			val.BalanceQty = Product.BalanceQty
			val.DiffQty = val.Qty - Product.BalanceQty
			if StockCount.FlagTemp == 0 {
				val.Active = true
				val.Remark = ""
			} else {
				val.Active = false
				val.Remark = "รอการปรับปรุง"
			}
			val.DocDate = StockCount.DocDate
			val.AverageCost = val.Price
			fullDataSub = append(fullDataSub, val)
		}
	}
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&StockCount)
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

//UpdateStockCount _
func UpdateStockCount(StockCount StockCount, user Users) (retID int64, errRet error) {
	docCheck, _ := GetStockCount(StockCount.ID)
	if docCheck.ID == 0 {
		errRet = errors.New("ไม่พบข้อมูล")
	}
	StockCount.Creator = docCheck.Creator
	StockCount.CreatedAt = docCheck.CreatedAt
	StockCount.CreditDay = docCheck.CreditDay
	StockCount.CreditDate = docCheck.CreditDate
	StockCount.EditedAt = time.Now()
	StockCount.Editor = &user
	StockCount.Active = docCheck.Active
	var fullDataSub []StockCountSub
	for _, val := range StockCount.StockCountSub {
		if val.Product.ID != 0 {
			Product, _ := GetProduct(val.Product.ID)
			val.CreatedAt = time.Now()
			val.Creator = &user
			val.DocNo = StockCount.DocNo
			val.Flag = StockCount.Flag
			val.BalanceQty = Product.BalanceQty
			val.DiffQty = val.Qty - Product.BalanceQty
			if StockCount.FlagTemp == 0 {
				val.Active = true
				val.Remark = ""
			} else {
				val.Active = false
				val.Remark = "รอการปรับปรุง"
			}
			val.AverageCost = val.Price
			val.DocDate = StockCount.DocDate
			fullDataSub = append(fullDataSub, val)
		}
	}
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Update(&StockCount)
	if err == nil {
		_, err = o.QueryTable("stock_count_sub").Filter("doc_no", StockCount.DocNo).Delete()
	}
	if err == nil {
		_, err = o.InsertMulti(len(fullDataSub), fullDataSub)
	}
	if err == nil {
		o.Commit()
		retID = id
	} else {
		o.Rollback()
	}
	errRet = err
	return retID, errRet
}

//GetStockCount _
func GetStockCount(ID int) (doc *StockCount, errRet error) {
	StockCount := &StockCount{}
	o := orm.NewOrm()
	o.QueryTable("stock_count").Filter("ID", ID).RelatedSel().One(StockCount)
	o.QueryTable("stock_count_sub").Filter("doc_no", StockCount.DocNo).RelatedSel().All(&StockCount.StockCountSub)
	doc = StockCount
	return doc, errRet
}

//GetStockCountList _
func GetStockCountList(term string, limit int, dateBegin, dateEnd string) (sup *[]StockCount, rowCount int, errRet error) {
	StockCount := &[]StockCount{}
	o := orm.NewOrm()
	qs := o.QueryTable("stock_count")
	condSub1 := orm.NewCondition()
	condSub2 := orm.NewCondition()
	cond1 := condSub1.And("doc_date__gte", dateBegin).And("doc_date__lte", dateEnd)
	qs = qs.SetCond(cond1)
	if dateBegin != "" && dateEnd != "" {
		cond2 := condSub2.Or("Member__Name__icontains", term).Or("DocNo__icontains", term).Or("Remark__icontains", term)
		cond1 = cond1.AndCond(cond2)
		qs = qs.SetCond(cond1)
	}
	qs.RelatedSel().Limit(limit).All(StockCount)
	return StockCount, len(*StockCount), errRet
}

//UpdateCancelStockCount _
func UpdateCancelStockCount(ID int, remark string, user Users) (retID int64, errRet error) {
	docCheck := &StockCount{}
	o := orm.NewOrm()
	o.QueryTable("stock_count").Filter("ID", ID).RelatedSel().One(docCheck)
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
		_, err = o.Raw("update stock_count_sub set active = false where doc_no = ?", docCheck.DocNo).Exec()
	}
	if err != nil {
		o.Rollback()
	} else {
		o.Commit()
	}
	errRet = err
	return retID, errRet
}

//UpdateActiveStockCount _
func UpdateActiveStockCount(ID int, user Users) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	orm.Debug = true
	_, err := o.Raw("update stock_count set active = true,flag_temp = 0,editor_id = ?,edited_at = now() where i_d = ?", user.ID, ID).Exec()
	if err != nil {
		o.Rollback()
	} else {
		_, err := o.Raw("update stock_count_sub set active = true where doc_no = (select stock_count.doc_no from stock_count where stock_count.i_d = ? limit 1)", ID).Exec()
		if err != nil {
			o.Rollback()
		} else {
			o.Commit()
		}
	}
	errRet = err
	return retID, errRet
}
