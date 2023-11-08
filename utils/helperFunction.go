package utils

import (
	"strconv"
	"time"
)

func GetTodaysDateToString() string {
	tm := time.Now()
	month := tm.Month().String()
	day := strconv.Itoa(tm.Day())
	yr := strconv.Itoa(tm.Year())
	todaysDate := day + "-" + month + "-" + yr

	return todaysDate
}
