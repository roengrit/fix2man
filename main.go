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
	beego.Router("/normal/list", &c.EntitryController{}, "post:ListEntity")
	beego.Router("/normal/add", &c.EntitryController{}, "get:NewEntity")
	beego.Router("/normal/update", &c.EntitryController{}, "post:UpdateEntity")

	beego.Router("/service/entitylist/json", &c.ServiceController{}, "get:ListEntityJSON")
	beego.Router("/service/userlist/json", &c.ServiceController{}, "get:GetUserListJSON")
	beego.Router("/service/user/json", &c.ServiceController{}, "get:GetUserJSON")
	beego.Router("/service/entitylist-p/json", &c.ServiceController{}, "get:ListEntityWithParentJSON")

	beego.Router("/create-request", &c.ReqController{})
	beego.Router("/request/read", &c.ReqController{}, "get:Read")
	beego.Router("/request/list", &c.ReqController{}, "get:ReqList;post:GetReqList")
	beego.Router("/request/change-status", &c.ReqController{}, "get:ChangeStatus;post:UpdateStatus")

	beego.Router("/supplier/list", &c.SupplierController{}, "get:SupList;post:GetSupList")
	beego.Router("/create-supplier", &c.SupplierController{}, "get:CreateSup;post:UpdateSup")
	beego.Router("/supplier/read", &c.SupplierController{}, "get:CreateSup")

	beego.Router("/create-receive", &c.RecController{})
	//beego.Router("/receive/read", &c.RecController{}, "get:Read")
	beego.Router("/receive/list", &c.RecController{}, "get:RecList;post:GetRecList")

	beego.Run()
}
