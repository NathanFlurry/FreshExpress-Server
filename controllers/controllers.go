package controllers

import "github.com/astaxie/beego"

func InitRouters() {
	// Route the controllers
	beego.Router("/buses", &BusesController{})
	beego.Router("/schedule", &ScheduleController{})
	beego.Router("/foods", &FoodsController{})
	beego.Router("/stops", &StopsController{})
}