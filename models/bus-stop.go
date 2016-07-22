package models

type BusStop struct {
	Id int
	Schedule []*ScheduleItem `orm:"reverse(many)"`
	LocationName string
	Address string
}
