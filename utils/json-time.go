package utils

import (
	"github.com/astaxie/beego/orm"
	"time"
	"encoding/json"
)

type JSONTime orm.DateTimeField

// Convert the time to the proper format
func (this JSONTime) MarshalJSON() ([]byte, error) {
	seconds := time.Time(this).Unix() // Find the number of seconds
	return json.Marshal(seconds)
}