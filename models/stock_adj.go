package models

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//StockAdj _
type StockAdj struct {
	ID      int
	Flag    int      //0 not process //1 in process
	Product *Product `orm:"rel(fk)"`
}

//StockOnly _
type StockOnly struct {
	Qty float64
}

//PreAllStockAdj _
type PreAllStockAdj struct {
	CreatedAt   time.Time
	ProductID   int
	BalanceQty  float64
	Qty         float64
	Price       float64
	AverageCost float64
	Flag        int
	DocType     int
	ID          int
	Tb          string
	ProductType int
}

func init() {
	orm.RegisterModel(new(StockAdj)) // Need to register model in init
}

//GetAllTransToProcessAvg _
func GetAllTransToProcessAvg(productID int) (num int64, PreAllStockAdj []PreAllStockAdj, err error) {
	var sql = `select all_table.* ,product.product_type from (
						select receive.created_at ,product_id as product_i_d ,0 as balance_qty,qty,price,average_cost,receive."flag",doc_type,receive_sub.i_d,'receive_sub' as tb   
						from receive_sub join receive on receive_sub.doc_no = receive.doc_no where receive.active and  receive_sub.active and product_id = ? 
				union all 
						select "pick_up".created_at ,product_id as product_i_d ,0 as balance_qty ,qty,price,average_cost,"pick_up"."flag",doc_type,pick_up_sub.i_d,'pick_up_sub' as tb   
						from pick_up_sub join "pick_up" on pick_up_sub.doc_no = "pick_up".doc_no where "pick_up".active and  pick_up_sub.active and product_id = ?  				 
				) as all_table JOIN product on all_table.product_i_d = product.i_d
				ORDER BY all_table.created_at,doc_type,all_table.i_d`
	sql = strings.Replace(sql, "?", strconv.Itoa(productID), -1)
	o := orm.NewOrm()
	num, err = o.Raw(sql).QueryRows(&PreAllStockAdj)
	return num, PreAllStockAdj, err
}

//GetAllTransToProcessSTK -
func GetAllTransToProcessSTK(productID int) {
	var sql = `select sum(qty) qty from (
						select qty    
						from receive_sub join receive on receive_sub.doc_no = receive.doc_no where receive.active and  receive_sub.active and product_id = ? 
				union all 
						select   qty  * -1 as qty 
						from pick_up_sub join "pick_up" on pick_up_sub.doc_no = "pick_up".doc_no where "pick_up".active and  pick_up_sub.active and product_id = ?  				 
				) as all_table   `
	sql = strings.Replace(sql, "?", strconv.Itoa(productID), -1)
	o := orm.NewOrm()
	var res StockOnly
	_ = o.Raw(sql).QueryRow(&res)
	_, _ = o.Raw("update product set balance_qty = ? , average_cost = 0 where i_d = ?", res.Qty, productID).Exec()
}

//CalAllAvgTrans _
func CalAllAvgTrans(productID int, updateTrans bool) (err error) {
	o := orm.NewOrm()
	num, list, err := GetAllTransToProcessAvg(productID)

	if num == 0 {
		return nil
	}
	if err != nil {
		return err
	}

	flagFirst := true
	var gAverageCost, gQty float64

	for _, val := range list {
		if flagFirst && (val.Flag == 2 || val.Flag == 3) {
			return errors.New("เกิดการขาย/เบิก ก่อนรับสต๊อคเข้าในระบบ")
		}
		if flagFirst && val.Flag == 4 && val.Qty-val.BalanceQty < 0 {
			return errors.New("เกิดการปรับปรุงสต๊อค (-) ก่อนรับสต๊อคเข้าในระบบ")
		}
		switch val.Flag {
		case 1:
			if !flagFirst {
				gAverageCost = ((val.Qty * val.Price) + (gQty * gAverageCost)) / (val.Qty + gQty)
				gQty = gQty + val.Qty
			} else {
				val, gAverageCost, gQty = firstReceive(gAverageCost, gQty, val)
			}
			if updateTrans {
				_, err = o.Raw("update "+val.Tb+" set average_cost = ? where i_d = ?", gAverageCost, val.ID).Exec()
			}
		case 2:
			gQty = gQty - val.Qty
			val.AverageCost = gAverageCost
			if updateTrans {
				_, err = o.Raw("update "+val.Tb+" set average_cost = ?,price = ? where i_d = ?", gAverageCost, gAverageCost, val.ID).Exec()
			}
			if gQty == 0 {
				gAverageCost = 0
			}
		case 3:
			gQty = gQty - val.Qty
			val.AverageCost = gAverageCost
			if updateTrans {
				_, err = o.Raw("update "+val.Tb+" set average_cost = ?,price = ? where i_d = ?", gAverageCost, gAverageCost, val.ID).Exec()
			}
			if gQty == 0 {
				gAverageCost = 0
			}
		case 4:
			if !flagFirst {
				if val.Qty-val.BalanceQty >= 0 {
					if val.Qty-val.BalanceQty == 0 {
						gQty = 0
						gAverageCost = val.AverageCost
					} else {
						gQty = val.Qty
						gAverageCost = val.AverageCost
					}
				} else {
					gQty = val.Qty
					gAverageCost = val.AverageCost
				}
			} else {
				val, gAverageCost, gQty = firstCountStock(gAverageCost, gQty, val)
			}
			if updateTrans {
				_, err = o.Raw("update "+val.Tb+" set average_cost = ? where i_d = ?", gAverageCost, val.ID).Exec()
			}
			if gQty == 0 {
				gAverageCost = 0
			}
		}
		flagFirst = false
		if err != nil {
			o.Rollback()
			return err
		}
	}
	_, err = o.Raw("update product set balance_qty = ? , average_cost = ? where i_d = ?", gQty, gAverageCost, productID).Exec()
	if err != nil {
		o.Rollback()
	} else {
		o.Commit()
	}
	return err
}

func firstReceive(AverageCost, Qty float64, tr PreAllStockAdj) (PreAllStockAdj, float64, float64) {
	tr.AverageCost = tr.Price
	AverageCost = tr.Price
	Qty = tr.Qty
	return tr, AverageCost, Qty
}

func firstCountStock(AverageCost, Qty float64, tr PreAllStockAdj) (PreAllStockAdj, float64, float64) {
	AverageCost = tr.AverageCost
	Qty = tr.Qty
	return tr, AverageCost, Qty
}

//CalAllAvg _
func CalAllAvg() {
	StockAdj := &[]StockAdj{}
	o := orm.NewOrm()
	qs := o.QueryTable("stock_adj")
	qs.Filter("flag", 1).Limit(1).Distinct().RelatedSel().All(StockAdj)
	if len(*StockAdj) > 0 {
		return
	}
	qs.Filter("flag", 0).Limit(100).RelatedSel().All(StockAdj)
	if len(*StockAdj) >= 1 {
		for _, val := range *StockAdj {
			_, _ = o.Raw("update stock_adj set flag = 1 where product_id = ?", val.Product.ID).Exec()
			o.Commit()
		}
	}
	qs.Filter("flag", 1).RelatedSel().All(StockAdj)
	if len(*StockAdj) >= 1 {
		for _, val := range *StockAdj {
			if val.Product.ProductType == 0 && !val.Product.Serial {
				CalAllAvgTrans(val.Product.ID, true)
			} else {
				GetAllTransToProcessSTK(val.Product.ID)
			}
			_, _ = o.Raw("delete from stock_adj where flag = 1 and product_id = ?", val.Product.ID).Exec()
			o.Commit()
		}
	}
}
