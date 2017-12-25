package main

import (
	"fix2man/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

//MainController _
type MainController struct {
	beego.Controller
}

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "host=localhost port=5432 user=postgres password=P@ssw0rd dbname=fixman sslmode=disable")
}

//Get _
func (this *MainController) Get() {
	name := "default"

	// Drop table and re-create.
	force := false

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	o := orm.NewOrm()
	o.Using("default") // Using default, you can use other database

	user := new(models.Users)
	user.Username = "slene"

	fmt.Println(o.Insert(user))
	this.Ctx.WriteString("hello world")
}

func main() {

	beego.Router("/", &MainController{})
	beego.Run()
}
