package models

import "github.com/astaxie/beego/orm"

type ScheduleItem struct {
	Id int
	Bus *Bus `orm:"rel(fk)"`
	Stop *BusStop `orm:"null;rel(fk)"`
	StartDate orm.DateTimeField
	EndDate orm.DateTimeField
}
