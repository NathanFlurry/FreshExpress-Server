package models

type Bus struct {
	Id int
	Foods []*FoodItem `orm:"rel(m2m)"`
	Schedule []*ScheduleItem `orm:"reverse(many)"`
}