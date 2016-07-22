package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type BusController struct {
	beego.Controller
}

func (this *BusController) Prepare() {

}

func (this *BusController) Get() {
	fmt.Println("Get")
	myStruct := map[string]interface{}{
		"hi": 5,
		"hello": "there",
	}
	this.Data["json"] = &myStruct
	this.ServeJSON()
}

