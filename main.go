package main

import (
	c "fix2man/controllers"

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

	beego.Router("/entity/location/depart/list", &c.LocationController{}, "get:GetDepartList")
	beego.Router("/entity/location/depart", &c.LocationController{}, "get:CreateDepart;post:UpdateDepart")

	//beego.Router("/entity/location/building/list", &c.LocationController{}, "get:GetBuildingList")
	//beego.Router("/entity/location/building", &c.LocationController{}, "get:CreateBuilding;post:UpdateBuilding")

	//beego.Router("/entity/location/class/list", &c.LocationController{}, "get:GetClassList")
	//beego.Router("/entity/location/class", &c.LocationController{}, "get:CreateClass;post:UpdateClass")

	//beego.Router("/entity/location/room/list", &c.LocationController{}, "get:GetRoomList")
	//beego.Router("/entity/location/room", &c.LocationController{}, "get:CreateRoom;post:UpdateRoom")

	beego.Router("/service/entity-list/json", &c.ServiceController{}, "get:ListEntityJSON")
	beego.Router("/service/entity-list-p/json", &c.ServiceController{}, "get:ListEntityWithParentJSON")

	beego.Router("/service/user/json", &c.ServiceController{}, "get:GetUserJSON")
	beego.Router("/service/user-list/json", &c.ServiceController{}, "get:GetUserListJSON")

	beego.Router("/create-request", &c.ReqController{})
	beego.Router("/request/read", &c.ReqController{}, "get:ReadReq")
	beego.Router("/request/change-status", &c.ReqController{}, "get:ChangeStatus;post:UpdateStatus")
	beego.Router("/request/list", &c.ReqController{}, "get:ReqList;post:GetReqList")

	beego.Router("/create-supplier", &c.SupplierController{}, "get:CreateSup;post:UpdateSup")
	beego.Router("/supplier/read", &c.SupplierController{}, "get:CreateSup")
	beego.Router("/supplier/list", &c.SupplierController{}, "get:SupList;post:GetSupList")

	beego.Router("/create-receive", &c.RecController{})
	beego.Router("/receive/list", &c.RecController{}, "get:RecList;post:GetRecList")

	beego.Run()
}
