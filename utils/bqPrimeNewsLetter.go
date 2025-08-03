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
	//OLD_URL: works in 2025 for certain days
	baseString := `https://www.ndtvprofit.com/business/stock-market-today-all-you-need-to-know-going-into-trade-on-`

	//till june 2025: works in 2025 for certain days
	baseString_2024 := `https://www.ndtvprofit.com/markets/stock-market-today-all-you-need-to-know-before-going-into-trade-on-`

	tm := time.Now()
	month := tm.Month().String()
	monthLowerCase := strings.ToLower(month)
	date := tm.Day()
	yr := YearToEdgePasser(tm.Year())
	yearAsString := strconv.Itoa(tm.Year())
	// monthShortner := getMonthAsPerCurrentMonth(month) // stubbing it 2025 aug since short month may not be in use

	//base string
	fullString := baseString + monthLowerCase + "-" + strconv.Itoa(date)
	fullString2024 := baseString_2024 + month + "-" + strconv.Itoa(date)

	//creating url
	fullStringUrl2024 := fullString2024 + "-" + strconv.Itoa(yr) //example: https://www.ndtvprofit.com/markets/stock-market-today-all-you-need-to-know-going-into-trade-on-july-23-2
	oneDecrementUrl2024 := fullString2024 + strconv.Itoa(yr-1)

	//old url decreament
	fullStringUrl := fullString + "-" + yearAsString //https://www.ndtvprofit.com/markets/stock-market-today-all-you-need-to-know-going-into-trade-on-july-30-2025
	//oneDecrementUrl := fullString + "-" + strconv.Itoa(yr-1)

	fmt.Println("Utils:GetBQPrime:fullstring2024: ", fullString2024)
	fmt.Println("Utils:GetBQPrime:fullstring: ", fullString)
	fmt.Println("Utils:GetBQPrime:fullstringURL: ", fullStringUrl)

	req, _ := http.Get(fullStringUrl)

	if req.StatusCode == 200 {
		fmt.Printf("in full string url")
		return fullStringUrl
	}

	req2024, _ := http.Get(fullStringUrl2024)

	if req2024.StatusCode == 200 {
		fmt.Printf("in 2024 full string url")
		return fullStringUrl2024
	}

	req2024Decr, _ := http.Get(oneDecrementUrl2024)

	if req2024Decr.StatusCode == 200 {
		fmt.Println("in 2024 one string decreament ")
		return oneDecrementUrl2024
	}

	//commenting this code as of 2025-08-03  since it not required
	// reqt, _ := http.Get(oneDecrementUrl)

	// if reqt.StatusCode == 200 {
	// 	fmt.Printf("in full string url")
	// 	return oneDecrementUrl
	// }

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
	if year == 2025 {
		return 3
	} else {
		return 0
	}
}
