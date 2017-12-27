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
	beego.Router("/prof-name", &c.GetNameController{})
	beego.Router("/forget-password", &c.ForgetController{})
	beego.Router("/normal", &c.EntitryController{})
	beego.Router("/normal/list", &c.EntitryController{}, "post:ListEntity")
	beego.Router("/normal/:id:int", &c.EntitryController{}, "get:GetEntity;put:UpdateEntity")
	beego.Run()
}
