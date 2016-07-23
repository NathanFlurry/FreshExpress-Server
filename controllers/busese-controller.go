package controllers

import (
	"github.com/astaxie/beego"
	"fresh-express-server/models"
)

type BusesController struct {
	beego.Controller
}

func (this *BusesController) Get() {
	qs := models.Ormer.QueryTable(new(models.Bus))

	var buses []*models.Bus
	_, _ = qs.RelatedSel().All(&buses)

	this.Data["json"] = map[string]interface{} {
		"buses": buses,
	}

	this.ServeJSON()
}

