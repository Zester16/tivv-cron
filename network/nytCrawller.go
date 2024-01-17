package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"stockpull/utils"
)

// calls news letter
func PostCrawlGetNYTimeArrayEveningBriefing() ([]utils.NewsLetterNytStruct, error) {

	r, err := http.NewRequest("POST", os.Getenv("nyt_cron"), nil)

	if err != nil {
		fmt.Println("network-post-PostCrawlGetNYTimeArrayEveningBriefing", err)
		return []utils.NewsLetterNytStruct{}, err
	}

	r.Header.Set("news_url", "https://www.nytimes.com/column/dealbook-newsletter")

	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
		fmt.Println("network-post-PostCrawlGetNYTimeArrayEveningBriefing", err)
		return []utils.NewsLetterNytStruct{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("network-post-PostCrawlGetNYTimeArrayEveningBriefing", err)
		return []utils.NewsLetterNytStruct{}, err
	}

	var newsLetterStruct []utils.NewsLetterNytStruct

	sb := string(body)
	fmt.Println(sb)
	json.Unmarshal([]byte(sb), &newsLetterStruct)
	fmt.Println("newsletterstruct", newsLetterStruct)
	return newsLetterStruct, nil

}
