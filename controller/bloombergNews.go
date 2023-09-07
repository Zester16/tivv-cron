package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetBloombergTechNews(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	//bloomberg := "https://www.bloomberg.com/newsletters/five-things-europe/latest"
	bq := "https://www.bqprime.com/markets/stock-market-today-all-you-need-to-know-going-into-trade-on-aug-07-02"
	req, _ := http.NewRequest("GET", bq, nil)
	//req.Header = http.Header{"Cookie": {"table: 0x7f0fe1d04bf8; exp_pref=AMER; seen_uk=1; country_code=US"}, "Host": {"www.bloomberg.com"}, "User-Agent": {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36"}}
	//req.Header.Set("User-Access", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36")
	//req.Header.Set("Cookie", "table: 0x7f0fe1d04bf8; exp_pref=AMER; seen_uk=1; country_code=US")
	rep, _ := client.Do(req)

	//body, err := http.Get(bloomberg)

	fmt.Print(rep.StatusCode)
	response, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		//log.Fatal(err)
		//w.Write([]byte("Not our page"))
	}
	rep.Body.Close()
	w.Write(response)

}

// gets bloomberg news
func GetBQPrimeTodaysAllYouNeedToKnowNews(w http.ResponseWriter, r *http.Request) {

	//fmt.Print(fullString)
	w.Write([]byte(GetBQPrimeUrl()))
}

func GetBQPrimeUrl() string {
	baseString := `https://www.bqprime.com/business/stock-market-today-all-you-need-to-know-going-into-trade-on-`
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
