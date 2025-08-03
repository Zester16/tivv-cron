package services

import (
	"errors"
	"fmt"
	"stockpull/network"
	"strconv"
	"strings"
	"time"
)

// tries to get particular day's url, by creating and checking url
func GetLivemintTopOfTheDayUrl() (string, error) {

	//example of url to constructhttps://www.livemint.com/mint-top-newsletter/minttopofthemorning02082025.html
	baseUrl := "https://www.livemint.com/mint-top-newsletter/minttopofthemorning{}.html"

	tm := time.Now()

	day := dateConverter(tm.Day())
	month := dateConverter(int(tm.Month()))
	year := tm.Year()

	fullDate := day + month + strconv.Itoa(year)

	fmt.Println("GetLivemintTopOfTheDayUrl date", fullDate)
	fullUrl := strings.Replace(baseUrl, "{}", fullDate, 1)

	status := network.CheckLiveMintNewsletterUrl(fullUrl)

	if status {
		return fullUrl, nil
	} else {
		return fullUrl, errors.New("Url does not exist")
	}

}

// checks if date or a day is less than 10, if yes it adds 0 before date and returns
// ex if input is 3 output i 03 as string
func dateConverter(input int) string {
	if input < 10 {
		return "0" + strconv.Itoa(input)
	}

	return strconv.Itoa(input)
}
