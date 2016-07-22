package controllers

import (
	"github.com/astaxie/beego"
	"fresh-express-server/models"
)

type StopController struct {
	beego.Controller
}

func (this *StopController) Get() {
	qs := (*models.GetOrmer()).QueryTable(new(models.BusStop))

	var stops []*models.ScheduleItem
	_, _ = qs.All(&stops)

	this.Data["json"] = map[string]interface{} {
		"stops": stops,
	}

	this.ServeJSON()
}