package utils

import (
	"fmt"
	"os"
)

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

	urlArrays := map[string]string{"dealbook": os.Getenv("nyt_dealbook")}

	return urlArrays[key]

}
