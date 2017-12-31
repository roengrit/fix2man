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
	beego.Router("/auth", &c.AuthController{})
	beego.Router("/logout", &c.LogoutController{})
	beego.Router("/prof-name", &c.GetNameController{}, "get:GetName")
	beego.Router("/forget-password", &c.ForgetController{})

	beego.Router("/normal", &c.EntitryController{})
	beego.Router("/normal/list", &c.EntitryController{}, "post:ListEntity")
	beego.Router("/normal/add", &c.EntitryController{}, "get:NewEntity")
	beego.Router("/normal/max", &c.EntitryController{}, "get:MaxEntity")
	beego.Router("/normal/update", &c.EntitryController{}, "post:UpdateEntity")

	beego.Router("/service/entitylist/json", &c.ServiceController{}, "get:ListEntityJson")
	beego.Router("/service/userlist/json", &c.ServiceController{}, "get:GetUserListJson")
	beego.Router("/service/user/json", &c.ServiceController{}, "get:GetUserJson")
	beego.Router("/service/entitylist-p/json", &c.ServiceController{}, "get:ListEntityWithParentJson")

	beego.Router("/create-request", &c.ReqController{})

	beego.Run()
}
