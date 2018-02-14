package main

import (
	c "fix2man/controllers"
	h "fix2man/helps"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "host=localhost port=5432 user=postgres password=P@ssw0rd dbname=fixman sslmode=disable")
}

func main() {

	name := "default"
	force := false                             // Drop table and re-create.
	verbose := true                            // Print log.
	err := orm.RunSyncdb(name, force, verbose) // Error.
	if err != nil {
		fmt.Println(err)
	}

	beego.Router("/", &c.AppController{})
	beego.Router("/auth", &c.UserController{})
	beego.Router("/change-pass", &c.UserController{}, "get:ChangePass;post:UpdatePass")
	beego.Router("/logout", &c.LogoutController{})
	beego.Router("/forget-password", &c.ForgetController{})

	beego.Router("/normal", &c.EntitryController{})
	beego.Router("/normal/add", &c.EntitryController{}, "get:NewEntity")
	beego.Router("/normal/update", &c.EntitryController{}, "post:UpdateEntity")
	beego.Router("/normal/list", &c.EntitryController{}, "post:ListEntity")

	beego.Router("/location/depart/list", &c.LocationController{}, "get:DepartList;post:GetDepartList")
	beego.Router("/location/depart", &c.LocationController{}, "get:CreateDepart;post:UpdateDepart")
	beego.Router("/location/depart/delete/?:id", &c.LocationController{}, "delete:DeleteDepart")

	beego.Router("/location/building/list", &c.LocationController{}, "get:BuildingList;post:GetBuildingList")
	beego.Router("/location/building", &c.LocationController{}, "get:CreateBuilding;post:UpdateBuilding")
	beego.Router("/location/building/delete/?:id", &c.LocationController{}, "delete:DeleteBuilding")

	beego.Router("/location/class/list", &c.LocationController{}, "get:ClassList;post:GetClassList")
	beego.Router("/location/class", &c.LocationController{}, "get:CreateClass;post:UpdateClass")
	beego.Router("/location/class/delete/?:id", &c.LocationController{}, "delete:DeleteClass")

	beego.Router("/location/room/list", &c.LocationController{}, "get:RoomList;post:GetRoomList")
	beego.Router("/location/room", &c.LocationController{}, "get:CreateRoom;post:UpdateRoom")
	beego.Router("/location/room/delete/?:id", &c.LocationController{}, "delete:DeleteRoom")

	beego.Router("/service/secure/json", &c.ServiceController{}, "get:GetXSRF")

	beego.Router("/service/entity/list/json", &c.ServiceController{}, "get:ListEntityJSON")
	beego.Router("/service/entity/list/p/json", &c.ServiceController{}, "get:ListEntityWithParentJSON")

	beego.Router("/service/user/json", &c.ServiceController{}, "get:GetUserJSON")
	beego.Router("/service/user/list/json", &c.ServiceController{}, "get:GetUserListJSON")
	beego.Router("/service/tech/list/json", &c.ServiceController{}, "get:GetTechListJSON")
	beego.Router("/service/avg", &c.ServiceNonAuthController{}, "get:CalItemAvg")
	beego.Router("/service/avg/o/?:id", &c.ServiceNonAuthController{}, "get:CalItemAvgByID")

	beego.Router("/member/?:id", &c.MemberController{}, "get:CreateMember;post:UpdateMember;delete:DeleteMember")
	beego.Router("/member/read/?:id", &c.MemberController{}, "get:CreateMember")
	beego.Router("/member/list", &c.MemberController{}, "get:MemberList;post:GetMemberList")

	beego.Router("/receive", &c.ReceiveController{})
	beego.Router("/receive/read", &c.ReceiveController{})
	beego.Router("/receive/cancel", &c.ReceiveController{}, "get:CancelReceive;post:UpdateCancelReceive")
	beego.Router("/receive/list", &c.ReceiveController{}, "get:ReceiveList;post:GetReceiveList")

	beego.Router("/pickup", &c.PickUpController{})
	beego.Router("/pickup/read", &c.PickUpController{})
	beego.Router("/pickup/cancel", &c.PickUpController{}, "get:CancelPickUp;post:UpdateCancelPickUp")
	beego.Router("/pickup/list", &c.PickUpController{}, "get:PickUpList;post:GetPickUpList")

	beego.Router("/create-request", &c.ReqController{})
	beego.Router("/request/read", &c.ReqController{}, "get:ReadReq")
	beego.Router("/request/change-status", &c.ReqController{}, "get:ChangeStatus;post:UpdateStatus")
	beego.Router("/request/list", &c.ReqController{}, "get:ReqList;post:GetReqList")

	beego.Router("/product/?:id", &c.ProductController{}, "get:CreateProduct;post:UpdateProduct;delete:DeleteProduct")
	beego.Router("/product/read/?:id", &c.ProductController{}, "get:CreateProduct")
	beego.Router("/product/list", &c.ProductController{}, "get:ProductList;post:GetProductList")
	beego.Router("/product/list/json", &c.ProductController{}, "get:ListProductJSON")
	beego.Router("/product/json", &c.ProductController{}, "get:GetProductJSON")
	beego.Router("/product/json/sn/avg", &c.ProductController{}, "get:GetProductSerialAvgJSON")

	beego.Router("/product-category/?:id", &c.ProductController{}, "get:CreateProductCate;post:UpdateProductCate;delete:DeleteProductCate")
	beego.Router("/product-category/list", &c.ProductController{}, "get:ProductCateList;post:GetProductCateList")

	beego.Router("/product-unit/?:id", &c.ProductController{}, "get:CreateProductUnit;post:UpdateProductUnit;delete:DeleteProductUnit")
	beego.Router("/product-unit/list", &c.ProductController{}, "get:ProductUnitList;post:GetProductUnitList")

	beego.Router("/assess", &c.AssessController{})

	beego.AddFuncMap("emptyDate", c.EmptyDateString)
	beego.AddFuncMap("ThCommaSep", h.ThCommaSep)
	beego.AddFuncMap("TextThCommaSep", h.TextThCommaSep)

	beego.Run()
}
