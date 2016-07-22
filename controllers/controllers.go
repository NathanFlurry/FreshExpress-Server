package controllers

import "github.com/astaxie/beego"

func init() {
	// Route the controllers
	beego.Router("/bus", BusController{})
	beego.Router("/schedule", ScheduleController{})
}