package controllers

import (
	"github.com/astaxie/beego"
	"fresh-express-server/models"
)

type BusController struct {
	beego.Controller
}

func (this *BusController) Get() {
	qs := (*models.GetOrmer()).QueryTable(new(models.Bus))

	var busses []*models.Bus
	_, _ = qs.All(&busses)

	this.Data["json"] = map[string]interface{} {
		"busses": busses,
	}

	this.ServeJSON()
}

