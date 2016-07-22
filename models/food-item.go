package models

type FoodItem struct {
	Id    int
	Buses []*Bus `orm:"reverse(many)"`
	Name  string
	Cost  float32
	// Nutrition facts?
}
