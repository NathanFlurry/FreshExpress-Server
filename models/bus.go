package models

type Bus struct {
	Id int
	Name string
	Foods []*FoodItem `orm:"rel(m2m)"`
	Schedule []*ScheduleItem `orm:"reverse(many)"`
}