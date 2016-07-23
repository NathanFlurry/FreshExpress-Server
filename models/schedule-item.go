package models

import (
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"fresh-express-server/utils"
)


type ScheduleItem struct {
	Id int
	Bus *Bus `orm:"rel(fk)"`
	Stop *BusStop `orm:"null;rel(fk)"`
	StartDate orm.DateTimeField
	EndDate orm.DateTimeField
}

func (this *ScheduleItem) MarshalJSON() ([]byte, error) {
	type Alias ScheduleItem
	return json.Marshal(&struct {
		*Alias
		StartDate utils.JSONTime
		EndDate utils.JSONTime
	}{
		Alias: (*Alias)(this),
		StartDate: utils.JSONTime(this.StartDate),
		EndDate: utils.JSONTime(this.EndDate),
	})
}