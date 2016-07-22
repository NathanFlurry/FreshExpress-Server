package main

import (
	"github.com/astaxie/beego"
	"fresh-express-server/models"
	"fresh-express-server/controllers"
)

func main() {
	// Init the ORM
	models.InitORM()

	// Init the routers
	controllers.InitRouters()

	// Run the server
	beego.Run()
}
