package main

import (
	"fresh-express-server/models"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql" // Import MySQL
	"fresh-express-server/controllers"
)

func init() {

}

func main() {
	// Register the models
	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.Bus))
	orm.RegisterModel(new(models.FoodItem))

	// Register the database
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root@/freshexpress") // user:password@/dbname

	//o := orm.NewOrm()

	// Routing
	beego.Router("/bus", &controllers.BusController{})
	beego.Router("/schedule", &controllers.ScheduleController{})

	// Run the server
	beego.Run()
}
