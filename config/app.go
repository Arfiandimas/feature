package config

import (
	"strconv"
	"time"
)

func GetCurrentYearAndMonth() string {
	year, month, _ := time.Now().Date()
	if int(month) < 10 {
		return "/" + strconv.Itoa(year) + "-" + "0" + strconv.Itoa(int(month))
	} else {
		return "/" + strconv.Itoa(year) + "-" + strconv.Itoa(int(month))
	}
}
