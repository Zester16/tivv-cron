package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// function checks if nqPrime newsletter is present for 2024, if it gives a status 404, then it will remove -3 with 2 and sets day
// if not then 0 is passed
func GetBQPrimeUrl() string {
	//OLD_URL
	baseString := `https://www.ndtvprofit.com/business/stock-market-today-all-you-need-to-know-going-into-trade-on-`
	//till june 2024
	baseString_2024 := `https://www.ndtvprofit.com/markets/stock-market-today-all-you-need-to-know-before-going-into-trade-on-`

	tm := time.Now()
	month := tm.Month().String()
	date := tm.Day()
	yr := YearToEdgePasser(tm.Year())                                                      // gives end trailing 1 2 etc
	fullString := baseString + getMonthAsPerCurrentMonth(month) + "-" + strconv.Itoa(date) //base string
	fullString2024 := baseString_2024 + getMonthAsPerCurrentMonth(month) + "-" + strconv.Itoa(date)
	fullStringUrl := fullString + "-" + strconv.Itoa(yr)
	oneDecrementUrl := fullString + "-" + strconv.Itoa(yr-1)

	fmt.Println("Utils:GetBQPrime:fullstring2024: ", fullString2024)
	fmt.Println("Utils:GetBQPrime:fullstring: ", fullString)
	fmt.Println("Utils:GetBQPrime:fullstringURL: ", fullStringUrl)

	req2024, _ := http.Get(fullString2024)

	if req2024.StatusCode == 200 {
		fmt.Printf("in 2024 full string url")
		return fullString2024
	}

	req, _ := http.Get(fullStringUrl)

	if req.StatusCode == 200 {
		fmt.Printf("in full string url")
		return fullStringUrl
	}

	reqt, _ := http.Get(oneDecrementUrl)

	if reqt.StatusCode == 200 {
		fmt.Printf("in full string url")
		return oneDecrementUrl
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
	if year == 2024 {
		return 3
	} else {
		return 0
	}
}
