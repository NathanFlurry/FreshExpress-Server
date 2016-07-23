package controllers

import (
	"github.com/astaxie/beego"
	"fresh-express-server/models"
)

type StopsController struct {
	beego.Controller
}

func (this *StopsController) Get() {
	qs := models.Ormer.QueryTable(new(models.BusStop))

	var stops []*models.BusStop
	_, _ = qs.RelatedSel().All(&stops)

	this.Data["json"] = map[string]interface{} {
		"items": stops,
	}

	this.ServeJSON()
}