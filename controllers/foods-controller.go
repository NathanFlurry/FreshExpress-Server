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

	var foods []*models.FoodItem
	_, _ = qs.RelatedSel().All(&foods)

	this.Data["json"] = map[string]interface{} {
		"foods": foods,
	}

	this.ServeJSON()
}
