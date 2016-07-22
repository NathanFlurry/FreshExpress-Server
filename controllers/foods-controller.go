package controllers

import (
	"github.com/astaxie/beego"
	"fresh-express-server/models"
)

type FoodsController struct {
	beego.Controller
}

func (this *FoodsController) Get() {
	qs := (*models.GetOrmer()).QueryTable(new(models.FoodItem))

	var foods []*models.ScheduleItem
	_, _ = qs.All(&foods)

	this.Data["json"] = map[string]interface{} {
		"foodsj": foods,
	}

	this.ServeJSON()
}
