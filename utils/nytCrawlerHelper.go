package utils

import (
	"fmt"
	"os"
)

var UrlArrays = map[string]string{"nyt_dealbook": os.Getenv("nyt_dealbook"), "nyt_morning_aus": os.Getenv("nyt_morning_aus"), "nyt_morning_apac": os.Getenv("nyt_morning_apac"), "nyt_morning_europe": os.Getenv("nyt_morning_europe"), "nyt_morning_us": os.Getenv("nyt_morning_us"), "nyt_evening_us": os.Getenv("nyt_evening_us")}

// re structures path of url
func NytArrayPopulator(news []NewsLetterNytStruct) []NewsLetterNytStruct {

	const NYT_URL = "https://www.nytimes.com"

	for i, n := range news {
		news[i].Url = NYT_URL + n.Url
	}
	fmt.Println("utils-NYTArrayPopulator")
	return news
}

func NytNewsArray(key string) string {

	urlArrays := map[string]string{"nyt_dealbook": os.Getenv("nyt_dealbook"), "nyt_morning_aus": os.Getenv("nyt_morning_aus"), "nyt_morning_apac": os.Getenv("nyt_morning_apac"), "nyt_morning_europe": os.Getenv("nyt_morning_europe"), "nyt_morning_us": os.Getenv("nyt_morning_us"), "nyt_evening_us": os.Getenv("nyt_evening_us")}
	return urlArrays[key]

}
