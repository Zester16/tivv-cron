package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// function checks if nqPrime newsletter is present for 2023, if it gives a status 404, then it will remove -2 and sets day
func GetBQPrimeUrl() string {
	baseString := `https://www.ndtvprofit.com/business/stock-market-today-all-you-need-to-know-going-into-trade-on-`
	tm := time.Now()
	month := tm.Month().String()
	date := tm.Day()
	yr := YearToEdgePasser(tm.Year())
	fullString := baseString + getMonthAsPerCurrentMonth(month) + "-" + strconv.Itoa(date)
	fullStringUrl := fullString + "-" + strconv.Itoa(yr)
	fmt.Printf(fullStringUrl)
	req, _ := http.Get(fullStringUrl)

	if req.StatusCode == 200 {
		fmt.Printf("in full string url")
		return fullStringUrl
	}

	return fullString

}
func getMonthAsPerCurrentMonth(month string) string {

	var mtn string
	if month == "January" {
		mtn = "jan"
	} else if month == "February" {
		mtn = "feb"
	} else if month == "August" {
		mtn = "aug"
	} else if month == "September" {
		mtn = "sept"
	} else if month == "October" {
		mtn = "oct"
	} else if month == "November" {
		mtn = "nov"
	} else if month == "December" {
		mtn = "dec"
	} else {
		mtn = strings.ToLower(month)
	}

	return mtn
}

func YearToEdgePasser(year int) int {
	if year == 2023 {
		return 2
	} else {
		return 0
	}
}
