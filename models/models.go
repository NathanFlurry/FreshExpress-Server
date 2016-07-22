package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

var ormer orm.Ormer

func init() {
	// Enable debug for the ORM
	orm.Debug = true

	// Register the models
	orm.RegisterModel(new(Bus), new(FoodItem), new(ScheduleItem), new(BusStop))

	// Register the database
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root@/freshexpress") // user:password@/dbname

	// Generate the DB
	err := orm.RunSyncdb("default", true, false)
	if err != nil {
		fmt.Println(err)
	}

	// Get the ormer
	ormer = orm.NewOrm()

	// Generate the items
	InitData()
}

func InitData() {
	// Create the bus
	bus := Bus{}

	// Create the foods
	foods := []FoodItem {
		{Buses: []*Bus{ &bus }, Name: "Banana", Cost: 10.00},
	}

	// Create the bus stops
	stops := []BusStop {
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
	schedule := make([]ScheduleItem, len(quickSchedule))
	for i, v := range quickSchedule {
		schedule[i] = ScheduleItem{
			Bus: &bus,
			Stop: &stops[v.SN],
			StartDate: orm.DateTimeField(time.Date(2016, time.Month(v.M), v.D, v.SH, v.SM, 0, 0, time.UTC)),
			EndDate: orm.DateTimeField(time.Date(2016, time.Month(v.M), v.D, v.EH, v.EM, 0, 0, time.UTC)),
		}
	}

	// Save the items
	//ormer.Insert(&bus) // TODO: Add back, getting divide by zero error
	for _, f := range foods {
		ormer.Insert(&f)
	}
	for _, s := range stops {
		ormer.Insert(&s)
	}
	for _, si := range schedule {
		ormer.Insert(&si)
	}
}

func GetOrmer() *orm.Ormer {
	return &ormer
}