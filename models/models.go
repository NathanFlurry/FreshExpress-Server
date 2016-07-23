package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var Ormer orm.Ormer

func InitORM() {
	// Enable debug for the ORM
	orm.Debug = false

	// Register the models
	orm.RegisterModel(new(Bus), new(BusStop), new(FoodItem), new(ScheduleItem))

	// Register the database
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root@/freshexpress") // user:password@/dbname

	// Generate the DB
	err := orm.RunSyncdb("default", true, false)
	if err != nil {
		fmt.Println(err)
	}

	// Get the ormer
	Ormer = orm.NewOrm()

	// Generate the items
	InitData()
}

func InitData() { // See http://www.discoverytriangle.org/cms2/wp-content/uploads/2016/07/07.2016_FE-Schedule2.pdf
	// Create the bus
	bus := Bus{Name: "Bus OG"}

	// Create the foods
	foods := []*FoodItem {
		{Name: "Broccoli", Description: "Rich in Vitamin K, Vitamin C, Fiber, Potassium, and Folate. Proven to prevent blood clots and even cancer.", Cost: 1.80},
		{Name: "Garlic", Description: "Garlic is great for preventing high blood presure and even cancer.", Cost: 2.10},
		{Name: "Spinach", Description: "Rich in vitamin A, Vitamin K, and folate. Spinach is especially good for your eyes.", Cost: 3.32},
		{Name: "Tomatoes", Description: "Even though tomatoes are not a vegetable, they are still great for preventing cardiovascular disease.", Cost: 4.20},

		{Name: "Blueberries", Description: "While mighty tasteful, blueberries are also very good at inhibiting radical cells in your body. It's great for preventing neurodegenerative diseases.", Cost: 2.56},
		{Name: "Orange", Description: "Citrus fruites are known for being rich in vitamin C which protects the body from free radicals.", Cost: 2.32},
		{Name: "Cranberries", Description: "Oh! Look! A patern! Apparently cranberries prevent cancer too.", Cost: 1.26},
		{Name: "Grapes", Description: "Used in wine. But alcohol causes cancer. So don't drink wine, just eat grapes.", Cost: 3.00},
		{Name: "Strawberries", Description: "They're red and delicious and not an apple. Who cares what superpowers they give you, just eat them because they're good.", Cost: 2.98},
	}

	// Create the bus stops
	stops := []*BusStop {
		{LocationName: "Brunson Lee/Educare", Address: "1300 N 48th St"},
		{LocationName: "Sidney P. Osborn", Address: "1720 E. Adams St"},
		{LocationName: "Downtown Phoenix", Address: "2nd Ave & Adams "},
		{LocationName: "Burton Barr Library", Address: "1221 N Central Ave"},
		{LocationName: "Sojourner Center", Address: "2330 E Fillmore Street"},

		{LocationName: "Mountain Park Health Center", Address: "3830 E Van Buren St"},
		{LocationName: "Apache Trails ASL", Address: "2428 E Apache Blvd"},
		{LocationName: "Tempe Public Library", Address: "3500 S Rural Road"},
		{LocationName: "Gateway CC- Central", Address: "1245 E Buckeye Road"},
		{LocationName: "Neighborhood Ministries", Address: "1918 W Van Buren St"},

		{LocationName: "Stepping Stone", Address: "1311 N 14th St"},
		{LocationName: "Washington Manor", Address: "1123 E Monroe "},
		{LocationName: "ASU Downtown Campus", Address: "1st St & Fillmore"},
		{LocationName: "Memorial Towers", Address: "1405 S 7th Ave"},
		{LocationName: "Phoenix Day", Address: "115 E Tonto St"},
	}

	// Create the schedule
	type QS struct { // Quick schedule item
		SN int // Stop num
		D int // Day
		M int // Month
		SH int // Start hour
		SM int // Start minute
		EH int // End hour
		EM int // End minute
	}

	quickSchedule := []QS {
		{SN: 0, D: 25, M: 7, SH: 7, SM: 30, EH: 9, EM: 00},
		{SN: 1, D: 25, M: 7, SH: 9, SM: 30, EH: 10, EM: 30},
		{SN: 2, D: 25, M: 7, SH: 11, SM: 30, EH: 12, EM: 30},
		{SN: 3, D: 25, M: 7, SH: 13, SM: 00, EH: 14, EM: 00},
		{SN: 4, D: 25, M: 7, SH: 14, SM: 30, EH: 15, EM: 30},

		{SN: 5, D: 26, M: 7, SH: 8, SM: 00, EH: 9, EM: 00},
		{SN: 6, D: 26, M: 7, SH: 9, SM: 30, EH: 10, EM: 30},
		{SN: 7, D: 26, M: 7, SH: 11, SM: 00, EH: 12, EM: 00},
		{SN: 8, D: 26, M: 7, SH: 13, SM: 00, EH: 14, EM: 00},
		{SN: 9, D: 26, M: 7, SH: 14, SM: 30, EH: 15, EM: 30},

		{SN: 10, D: 27, M: 7, SH: 9, SM: 30, EH: 10, EM: 30},
		{SN: 11, D: 27, M: 7, SH: 11, SM: 00, EH: 12, EM: 00},
		{SN: 12, D: 27, M: 7, SH: 13, SM: 00, EH: 14, EM: 30},
		{SN: 13, D: 27, M: 7, SH: 14, SM: 30, EH: 15, EM: 30},
		{SN: 14, D: 27, M: 7, SH: 18, SM: 00, EH: 19, EM: 00},
	}

	// Generate the schedule
	schedule := make([]*ScheduleItem, len(quickSchedule))
	for i, v := range quickSchedule {
		schedule[i] = &ScheduleItem{
			Bus: &bus,
			//Stop: stops[v.SN], // TODO: Figure out why stops is not working
			StartDate: orm.DateTimeField(time.Date(2016, time.Month(v.M), v.D, v.SH, v.SM, 0, 0, time.UTC)),
			EndDate: orm.DateTimeField(time.Date(2016, time.Month(v.M), v.D, v.EH, v.EM, 0, 0, time.UTC)),
		}
	}

	// Save the items
	Ormer.Insert(&bus)
	HandleInsertMulti(Ormer.InsertMulti(len(foods), foods))
	HandleInsertMulti(Ormer.InsertMulti(len(stops), stops))
	HandleInsertMulti(Ormer.InsertMulti(len(schedule), schedule))

	// HACK: Change the ID by hand because this is being a butt
	for i, _ := range schedule {
		Ormer.QueryTable(new(ScheduleItem)).Filter("Id", i + 1).Update(orm.Params{
			"stop_id": quickSchedule[i].SN + 1,
		})
	}
}

func HandleInsertMulti(successNum int64, err error) {
	if err != nil {
		fmt.Errorf("Success num: %v\n%v\n", successNum, err)
	} else {
		fmt.Printf("Successfully inserted %v\n", successNum)
	}
}
