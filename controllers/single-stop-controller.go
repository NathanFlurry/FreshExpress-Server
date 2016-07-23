package controllers

import (
	"github.com/astaxie/beego"
	"fresh-express-server/models"
)

type SingleStopController struct {
	beego.Controller
}

func (this *SingleStopController) Get() {
	id := this.Ctx.Input.Param(":id")

	var stop models.BusStop
	models.Ormer.QueryTable(new(models.BusStop)).RelatedSel().Filter("Id", id).One(&stop)
	models.Ormer.LoadRelated(&stop, "Schedule")

	this.Data["json"] = map[string]interface{} {
		"stop": stop,
	}

	this.ServeJSON()
}