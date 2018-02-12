package models

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Product _
type Product struct {
	ID              int
	Lock            bool
	Name            string           `orm:"size(300)"`
	Barcode         string           `orm:"size(13)"`
	AverageCost     float64          `orm:"digits(12);decimals(2)"`
	BalanceQty      float64          `orm:"digits(12);decimals(2)"`
	SalePrice       float64          `orm:"digits(12);decimals(2)"`
	Unit            *Unit            `orm:"rel(fk)"`
	ProductCategory *ProductCategory `orm:"rel(fk)"`
	ProductType     int
	ImagePath1      string `orm:"size(300)"`
	ImageBase64     string `orm:"-"`
	Remark          string `orm:"size(100)"`
	FixCost         bool
	Active          bool
	Serial          bool
	Creator         *Users    `orm:"rel(fk)"`
	CreatedAt       time.Time `orm:"auto_now_add;type(datetime)"`
	Editor          *Users    `orm:"null;rel(fk)"`
	EditedAt        time.Time `orm:"null;auto_now;type(datetime)"`
}

//ProductCategory _
type ProductCategory struct {
	ID        int
	Lock      bool
	Name      string `orm:"size(300)"`
	Active    bool
	Creator   *Users    `orm:"rel(fk)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Editor    *Users    `orm:"null;rel(fk)"`
	EditedAt  time.Time `orm:"null;auto_now;type(datetime)"`
}

//Unit _
type Unit struct {
	ID        int
	Lock      bool
	Name      string `orm:"size(300)"`
	Active    bool
	Creator   *Users    `orm:"rel(fk)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Editor    *Users    `orm:"null;rel(fk)"`
	EditedAt  time.Time `orm:"null;auto_now;type(datetime)"`
}

//PriceOnly _
type PriceOnly struct {
	Price float64
}

//AverageCostOnly _
type AverageCostOnly struct {
	AverageCost float64
}

func init() {
	orm.RegisterModel(new(Product), new(ProductCategory), new(Unit)) // Need to register model in init
}

//GetProductList _
func GetProductList(top int, term string) (num int64, productList []Product, err error) {
	var sql = `SELECT T0.i_d,T0.name,T0.lock, T1.i_d as unit_id,T1.name as unit_name
			   FROM product T0	
			   JOIN unit T1 ON T0.unit_id = T1.i_d		    
			   WHERE lower(T0.name) like lower(?) order by T0.name limit {0}`
	if top == 0 {
		sql = strings.Replace(sql, "limit {0}", "", -1)
	} else {
		sql = strings.Replace(sql, "{0}", strconv.Itoa(top), -1)
	}
	o := orm.NewOrm()
	num, err = o.Raw(sql, "%"+term+"%").QueryRows(&productList)
	return num, productList, err
}

//GetProductRawList _
func GetProductRawList(top int, term string) (num int64, productList []Product, err error) {
	var sql = `SELECT T0.i_d,T0.name,T0.lock, T1.i_d as unit_id,T1.name as unit_name
			   FROM product T0	
			   JOIN unit T1 ON T0.unit_id = T1.i_d		    
			   WHERE T0.product_type = 2  and lower(T0.name) like lower(?) order by T0.name limit {0}`
	if top == 0 {
		sql = strings.Replace(sql, "limit {0}", "", -1)
	} else {
		sql = strings.Replace(sql, "{0}", strconv.Itoa(top), -1)
	}
	o := orm.NewOrm()
	num, err = o.Raw(sql, "%"+term+"%").QueryRows(&productList)
	return num, productList, err
}

//GetManagmentProductList _
func GetManagmentProductList(term string, limit int) (pro *[]Product, rowCount int, errRet error) {
	productlist := &[]Product{}
	o := orm.NewOrm()
	qs := o.QueryTable("product")
	cond := orm.NewCondition()
	cond1 := cond.Or("Name__icontains", term).
		Or("Remark__icontains", term)
	qs.SetCond(cond1).RelatedSel().Limit(limit).All(productlist)
	return productlist, len(*productlist), errRet
}

//GetProductCateList _
func GetProductCateList(term string, limit int) (cate *[]ProductCategory, rowCount int, errRet error) {
	ProductCategory := &[]ProductCategory{}
	o := orm.NewOrm()
	qs := o.QueryTable("product_category")
	cond := orm.NewCondition()
	cond1 := cond.Or("Name__icontains", term)
	qs.SetCond(cond1).RelatedSel().Limit(limit).All(ProductCategory)
	return ProductCategory, len(*ProductCategory), errRet
}

//GetProductUnitList _
func GetProductUnitList(term string, limit int) (cate *[]Unit, rowCount int, errRet error) {
	Unit := &[]Unit{}
	o := orm.NewOrm()
	qs := o.QueryTable("unit")
	cond := orm.NewCondition()
	cond1 := cond.Or("Name__icontains", term)
	qs.SetCond(cond1).RelatedSel().Limit(limit).All(Unit)
	return Unit, len(*Unit), errRet
}

//GetProduct _
func GetProduct(ID int) (pro *Product, errRet error) {
	Product := &Product{}
	o := orm.NewOrm()
	o.QueryTable("product").Filter("ID", ID).RelatedSel().One(Product)
	return Product, errRet
}

//GetProductSerialAvg _
func GetProductSerialAvg(SN string) float64 {

	var sql = ` select price  as average_cost 
					from receive_sub 
					join receive on receive_sub.doc_no = receive.doc_no  
					where receive.active and  receive_sub.active  and receive_sub.serial = '{1}'
					order by receive.doc_date desc
				 `
	sql = strings.Replace(sql, "{1}", SN, -1)
	o := orm.NewOrm()
	var res AverageCostOnly
	_ = o.Raw(sql).QueryRow(&res)
	return res.AverageCost
}

//GetProductCate _
func GetProductCate(ID int) (cate *ProductCategory, errRet error) {
	ProductCategory := &ProductCategory{}
	o := orm.NewOrm()
	o.QueryTable("product_category").Filter("ID", ID).RelatedSel().One(ProductCategory)
	return ProductCategory, errRet
}

//GetProductUnit _
func GetProductUnit(ID int) (unit *Unit, errRet error) {
	Unit := &Unit{}
	o := orm.NewOrm()
	o.QueryTable("unit").Filter("ID", ID).RelatedSel().One(Unit)
	return Unit, errRet
}

//CreateProduct _
func CreateProduct(pro Product) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&pro)
	o.Commit()
	if err == nil {
		retID = id
	}
	return retID, err
}

//CreateProductCate _
func CreateProductCate(cate ProductCategory) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&cate)
	o.Commit()
	if err == nil {
		retID = id
	}
	return retID, err
}

//CreateProductUnit _
func CreateProductUnit(unit Unit) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&unit)
	o.Commit()
	if err == nil {
		retID = id
	}
	return retID, err
}

//UpdateProduct _
func UpdateProduct(pro Product, isNewImage bool) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetProduct(pro.ID)
	if getUpdate.Lock {
		errRet = errors.New("ข้อมูลถูก Lock ไม่สามารถแก้ไขได้")
	}
	if getUpdate == nil {
		errRet = errors.New("ไม่พบข้อมูล")
	} else if errRet == nil {
		if !isNewImage {
			pro.ImagePath1 = getUpdate.ImagePath1
		}
		pro.BalanceQty = getUpdate.BalanceQty
		if !pro.FixCost {
			pro.AverageCost = getUpdate.AverageCost
		}
		pro.CreatedAt = getUpdate.CreatedAt
		pro.Creator = getUpdate.Creator
		if num, errUpdate := o.Update(&pro); errUpdate != nil {
			errRet = errUpdate
			_ = num
		}
	}
	return errRet
}

//UpdateProductCate _
func UpdateProductCate(cate ProductCategory) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetProductCate(cate.ID)
	if getUpdate.Lock {
		errRet = errors.New("ข้อมูลถูก Lock ไม่สามารถแก้ไขได้")
	}
	if getUpdate == nil {
		errRet = errors.New("ไม่พบข้อมูล")
	} else if errRet == nil {
		cate.CreatedAt = getUpdate.CreatedAt
		cate.Creator = getUpdate.Creator
		if num, errUpdate := o.Update(&cate); errUpdate != nil {
			errRet = errUpdate
			_ = num
		}
	}
	return errRet
}

//UpdateProductUnit _
func UpdateProductUnit(unit Unit) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetProductUnit(unit.ID)
	if getUpdate.Lock {
		errRet = errors.New("ข้อมูลถูก Lock ไม่สามารถแก้ไขได้")
	}
	if getUpdate == nil {
		errRet = errors.New("ไม่พบข้อมูล")
	} else if errRet == nil {
		unit.CreatedAt = getUpdate.CreatedAt
		unit.Creator = getUpdate.Creator
		if num, errUpdate := o.Update(&unit); errUpdate != nil {
			errRet = errUpdate
			_ = num
		}
	}
	return errRet
}

//DeleteProduct _
func DeleteProduct(ID int) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetProduct(ID)
	if getUpdate.Lock {
		errRet = errors.New("ข้อมูลถูก Lock ไม่สามารถแก้ไขได้")
	}
	if num, errDelete := o.Delete(&Product{ID: ID}); errDelete != nil && errRet == nil {
		errRet = errDelete
		_ = num
	}
	return errRet
}

//DeleteProductCate _
func DeleteProductCate(ID int) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetProductCate(ID)
	if getUpdate.Lock {
		errRet = errors.New("ข้อมูลถูก Lock ไม่สามารถแก้ไขได้")
	}
	if num, errDelete := o.Delete(&ProductCategory{ID: ID}); errDelete != nil && errRet == nil {
		errRet = errDelete
		_ = num
	}
	return errRet
}

//DeleteProductUnit _
func DeleteProductUnit(ID int) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetProductUnit(ID)
	if getUpdate.Lock {
		errRet = errors.New("ข้อมูลถูก Lock ไม่สามารถแก้ไขได้")
	}
	if num, errDelete := o.Delete(&Unit{ID: ID}); errDelete != nil && errRet == nil {
		errRet = errDelete
		_ = num
	}
	return errRet
}

//GetAllProductCategory _
func GetAllProductCategory() (pro *[]ProductCategory) {
	ProductCategory := &[]ProductCategory{}
	o := orm.NewOrm()
	o.QueryTable("product_category").RelatedSel().All(ProductCategory)
	return ProductCategory
}

//GetAllProductUnit _
func GetAllProductUnit() (pro *[]Unit) {
	Unit := &[]Unit{}
	o := orm.NewOrm()
	o.QueryTable("unit").RelatedSel().All(Unit)
	return Unit
}
