package models

import "github.com/astaxie/beego/orm"

type ScheduleItem struct {
	Id int
	Bus *Bus `orm:"rel(fk)"`
	Stop *BusStop `orm:"rel(fk)"`
	StartDate orm.DateTimeField
	EndDate orm.DateTimeField
}
