package controllers

import (
	"github.com/astaxie/beego"
	"fresh-express-server/models"
)

type ScheduleController struct {
	beego.Controller
}

func (this *ScheduleController) Get() {
	qs := models.Ormer.QueryTable(new(models.ScheduleItem))

	var items []*models.ScheduleItem
	_, _ = qs.RelatedSel().All(&items)

	this.Data["json"] = map[string]interface{} {
		"items": items,
	}

	this.ServeJSON()
}