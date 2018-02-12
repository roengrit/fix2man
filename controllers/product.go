package controllers

import (
	h "fix2man/helps"
	m "fix2man/models"
	"fmt"
	"html/template"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/form"
	"github.com/google/uuid"
)

//ProductController _
type ProductController struct {
	BaseController
}

//ListProductJSON  _
func (c *ProductController) ListProductJSON() {
	term := strings.TrimSpace(c.GetString("query"))
	raw, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("raw"), 10, 32)
	ret := m.NormalModel{}
	var err error
	var rowCount int64
	var lists []m.Product
	fmt.Println(raw)
	if raw == 0 {
		rowCount, lists, err = m.GetProductList(15, term)
	} else {
		rowCount, lists, err = m.GetProductRawList(15, term)
	}
	if err == nil {
		ret.RetOK = true
		ret.RetCount = rowCount
		ret.ListData = lists
		if rowCount == 0 {
			ret.RetOK = false
			ret.RetData = "ไม่พบข้อมูล"
		}
	} else {
		ret.RetOK = false
		ret.RetData = "ไม่พบข้อมูล"
	}
	c.Data["json"] = ret
	c.ServeJSON()
}

//GetProductJSON  _
func (c *ProductController) GetProductJSON() {
	ID, _ := strconv.ParseInt(c.GetString("id"), 10, 32)
	ret := m.NormalModel{}
	product, err := m.GetProduct(int(ID))
	if err == nil {
		ret.RetOK = true
		ret.Data1 = product
	} else {
		ret.RetOK = false
		ret.RetData = "ไม่พบข้อมูล"
	}
	c.Data["json"] = ret
	c.ServeJSON()
}

//GetProductSerialAvgJSON  _
func (c *ProductController) GetProductSerialAvgJSON() {
	SN := c.GetString("sn")
	ret := m.NormalModel{}
	product := m.GetProductSerialAvg(SN)
	ret.RetOK = true
	ret.Data1 = product
	c.Data["json"] = ret
	c.ServeJSON()
}

//CreateProduct _
func (c *ProductController) CreateProduct() {
	proID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	if strings.Contains(c.Ctx.Request.URL.RequestURI(), "read") {
		c.Data["r"] = "readonly"
	}
	if proID == 0 {
		c.Data["title"] = "สร้าง สินค้า"
	} else {
		c.Data["title"] = "แก้ไข สินค้า"
		pro, _ := m.GetProduct(int(proID))
		if len(pro.ImagePath1) > 0 {
			base64, _ := h.File64Encode(pro.ImagePath1)
			pro.ImageBase64 = base64
		}
		c.Data["m"] = pro
	}
	c.Data["ret"] = m.NormalModel{}
	c.Data["ProductCategory"] = m.GetAllProductCategory()
	c.Data["Unit"] = m.GetAllProductUnit()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "product/product.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "product/product-script.html"
	c.LayoutSections["html_head"] = "product/product-style.html"
	c.Render()
}

//UpdateProduct _
func (c *ProductController) UpdateProduct() {

	var pro m.Product
	decoder := form.NewDecoder()
	err := decoder.Decode(&pro, c.Ctx.Request.Form)
	ret := m.NormalModel{}
	actionUser, _ := m.GetUser(h.GetUser(c.Ctx.Request))
	file, header, _ := c.GetFile("ImgProduct")
	isNewImage := false
	if file != nil {
		fileName := header.Filename
		fileName = uuid.New().String() + filepath.Ext(fileName)
		filePathSave := "data/product/" + fileName
		err = c.SaveToFile("ImgProduct", filePathSave)
		if err == nil {
			isNewImage = true
			pro.ImagePath1 = filePathSave
			base64, errBase64 := h.File64Encode(filePathSave)
			err = errBase64
			pro.ImageBase64 = base64
		}
	}
	_ = file

	ret.RetOK = true
	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else if c.GetString("Name") == "" {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุชื่อ"
	}

	if ret.RetOK && pro.ID == 0 {
		pro.CreatedAt = time.Now()
		pro.Creator = &actionUser
		ID, err := m.CreateProduct(pro)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			pro.ID = int(ID)
			ret.RetData = "บันทึกสำเร็จ"
		}
	} else if ret.RetOK && pro.ID > 0 {
		pro.EditedAt = time.Now()
		pro.Editor = &actionUser
		err := m.UpdateProduct(pro, isNewImage)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	}
	if pro.ID == 0 {
		c.Data["title"] = "สร้าง สินค้า"
		c.Data["m"] = pro
	} else {
		c.Data["title"] = "แก้ไข สินค้า"
		pro, _ := m.GetProduct(int(pro.ID))
		if len(pro.ImagePath1) > 0 {
			base64, _ := h.File64Encode(pro.ImagePath1)
			pro.ImageBase64 = base64
		}
		c.Data["m"] = pro
	}
	c.Data["ret"] = ret
	c.Data["ProductCategory"] = m.GetAllProductCategory()
	c.Data["Unit"] = m.GetAllProductUnit()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "product/product.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "product/product-script.html"
	c.Render()
}

//DeleteProduct _
func (c *ProductController) DeleteProduct() {
	ID, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	ret := m.NormalModel{}
	ret.RetOK = true
	err := m.DeleteProduct(int(ID))
	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else {
		ret.RetData = "ลบข้อมูลสำเร็จ"
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//ProductList _
func (c *ProductController) ProductList() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["title"] = "สินค้า"
	c.Layout = "layout.html"
	c.TplName = "product/product-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "product/product-list-script.html"
	c.Render()
}

//GetProductList _
func (c *ProductController) GetProductList() {
	ret := m.NormalModel{}
	ret.RetOK = true
	top, _ := strconv.ParseInt(c.GetString("top"), 10, 32)
	term := c.GetString("txt-search")
	lists, rowCount, err := m.GetManagmentProductList(term, int(top))
	if err == nil {
		ret.RetOK = true
		ret.RetCount = int64(rowCount)
		ret.RetData = h.GenProHTML(*lists)
		if rowCount == 0 {
			ret.RetData = h.HTMLProNotFoundRows
		}
	} else {
		ret.RetOK = false
		ret.RetData = strings.Replace(h.HTMLProError, "{err}", err.Error(), -1)
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//CreateProductCate _
func (c *ProductController) CreateProductCate() {
	cateID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	if strings.Contains(c.Ctx.Request.URL.RequestURI(), "read") {
		c.Data["r"] = "readonly"
	}
	if cateID == 0 {
		c.Data["title"] = "สร้าง หมวดสินค้า"
	} else {
		c.Data["title"] = "แก้ไข หมวดสินค้า"
		cate, _ := m.GetProductCate(int(cateID))
		c.Data["m"] = cate
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "product-category/product-cate.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "product-category/product-cate-script.html"
	c.Render()
}

//UpdateProductCate _
func (c *ProductController) UpdateProductCate() {

	var cate m.ProductCategory
	decoder := form.NewDecoder()
	err := decoder.Decode(&cate, c.Ctx.Request.Form)
	ret := m.NormalModel{}
	actionUser, _ := m.GetUser(h.GetUser(c.Ctx.Request))

	ret.RetOK = true
	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else if c.GetString("Name") == "" {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุชื่อ"
	}

	if ret.RetOK && cate.ID == 0 {
		cate.CreatedAt = time.Now()
		cate.Creator = &actionUser
		_, err := m.CreateProductCate(cate)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	} else if ret.RetOK && cate.ID > 0 {
		cate.EditedAt = time.Now()
		cate.Editor = &actionUser
		err := m.UpdateProductCate(cate)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//DeleteProductCate _
func (c *ProductController) DeleteProductCate() {
	ID, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	ret := m.NormalModel{}
	ret.RetOK = true
	err := m.DeleteProductCate(int(ID))
	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else {
		ret.RetData = "ลบข้อมูลสำเร็จ"
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//ProductCateList _
func (c *ProductController) ProductCateList() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["title"] = "หมวดสินค้า"
	c.Layout = "layout.html"
	c.TplName = "product-category/product-cate-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "product-category/product-cate-list-script.html"
	c.Render()
}

//GetProductCateList _
func (c *ProductController) GetProductCateList() {
	ret := m.NormalModel{}
	ret.RetOK = true
	top, _ := strconv.ParseInt(c.GetString("top"), 10, 32)
	term := c.GetString("txt-search")
	lists, rowCount, err := m.GetProductCateList(term, int(top))
	if err == nil {
		ret.RetOK = true
		ret.RetCount = int64(rowCount)
		ret.RetData = h.GenProCateHTML(*lists)
		if rowCount == 0 {
			ret.RetData = h.HTMLProCateNotFoundRows
		}
	} else {
		ret.RetOK = false
		ret.RetData = strings.Replace(h.HTMLProCateError, "{err}", err.Error(), -1)
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//CreateProductUnit _
func (c *ProductController) CreateProductUnit() {
	unitID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	if strings.Contains(c.Ctx.Request.URL.RequestURI(), "read") {
		c.Data["r"] = "readonly"
	}
	if unitID == 0 {
		c.Data["title"] = "สร้าง หน่วย"
	} else {
		c.Data["title"] = "แก้ไข หน่วย"
		unit, _ := m.GetProductUnit(int(unitID))
		c.Data["m"] = unit
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "product-unit/product-unit.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "product-unit/product-unit-script.html"
	c.Render()
}

//UpdateProductUnit _
func (c *ProductController) UpdateProductUnit() {

	var unit m.Unit
	decoder := form.NewDecoder()
	err := decoder.Decode(&unit, c.Ctx.Request.Form)
	ret := m.NormalModel{}
	actionUser, _ := m.GetUser(h.GetUser(c.Ctx.Request))

	ret.RetOK = true
	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else if c.GetString("Name") == "" {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุชื่อ"
	}

	if ret.RetOK && unit.ID == 0 {
		unit.CreatedAt = time.Now()
		unit.Creator = &actionUser
		_, err := m.CreateProductUnit(unit)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	} else if ret.RetOK && unit.ID > 0 {
		unit.EditedAt = time.Now()
		unit.Editor = &actionUser
		err := m.UpdateProductUnit(unit)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//DeleteProductUnit _
func (c *ProductController) DeleteProductUnit() {
	ID, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	ret := m.NormalModel{}
	ret.RetOK = true
	err := m.DeleteProductUnit(int(ID))
	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else {
		ret.RetData = "ลบข้อมูลสำเร็จ"
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//ProductUnitList _
func (c *ProductController) ProductUnitList() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["title"] = "หน่วย"
	c.Layout = "layout.html"
	c.TplName = "product-unit/product-unit-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "product-unit/product-unit-list-script.html"
	c.Render()
}

//GetProductUnitList _
func (c *ProductController) GetProductUnitList() {
	ret := m.NormalModel{}
	ret.RetOK = true
	top, _ := strconv.ParseInt(c.GetString("top"), 10, 32)
	term := c.GetString("txt-search")
	lists, rowCount, err := m.GetProductUnitList(term, int(top))
	if err == nil {
		ret.RetOK = true
		ret.RetCount = int64(rowCount)
		ret.RetData = h.GenProUnitHTML(*lists)
		if rowCount == 0 {
			ret.RetData = h.HTMLProUnitNotFoundRows
		}
	} else {
		ret.RetOK = false
		ret.RetData = strings.Replace(h.HTMLProUnitError, "{err}", err.Error(), -1)
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}
