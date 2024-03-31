package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"stockpull/model"
	"stockpull/utils"
)

// struct for sending proper header
type NYTNewsBodyStruct struct {
	NewsUrl string `json:"news_url"`
}

// calls news letter
func PostCrawlGetNYTimeArrayEveningBriefing(key string) ([]utils.NewsLetterNytStruct, error) {
	nytNewsArray := model.BlmTest.GetNYTUrls()
	url := nytNewsArray[key]

	newsUrl := NYTNewsBodyStruct{NewsUrl: url}

	newsUrlJSON, _ := json.Marshal(newsUrl)

	fmt.Println("", bytes.NewBuffer(newsUrlJSON))

	r, err := http.NewRequest("POST", os.Getenv("nyt_cron"), bytes.NewBuffer(newsUrlJSON))

	if err != nil {
		fmt.Println("network/post/PostCrawlGetNYTimeArray", err)
		return []utils.NewsLetterNytStruct{}, err
	}

	r.Header.Set("news_url", url)
	r.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
		fmt.Println("network/post/PostCrawlGetNYTimeArray http Error:", err)
		return []utils.NewsLetterNytStruct{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("network/post/PostCrawlGetNYTimeArray body Error:", err)
		return []utils.NewsLetterNytStruct{}, err
	}

	var newsLetterStruct []utils.NewsLetterNytStruct

	sb := string(body)
	fmt.Println(sb)

	json.Unmarshal([]byte(sb), &newsLetterStruct)

	fmt.Println("newsletterstruct", newsLetterStruct)

	return newsLetterStruct, nil

}
