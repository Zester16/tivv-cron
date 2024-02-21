package cronjobs

import (
	"encoding/json"
	"fmt"
	"stockpull/datasource"
	"stockpull/network"
	"stockpull/utils"
)

type NYTNewsCachingStruct struct {
	eventType        string
	newsLetterStruct []utils.NewsLetterNytStruct
}

func SetNYTNewsLetterToRedis() {
	rdb := datasource.RedisConnect()
	c := make(chan NYTNewsCachingStruct)
	for key, value := range utils.UrlArrays {
		fmt.Println(key, value)
		go getNewsAndSetData(key, c)
	}
	counter := 0
	for x := range c {
		fmt.Println("nytNewsletterCronJob", x.eventType)

		j, err := json.Marshal(x.newsLetterStruct)
		if err != nil {
			fmt.Println("LivemintNewsMarshal Error: ", err)
		}
		rdb.RedisDBConnector.Set(ctx, x.eventType, j, 0).Err()
		counter = counter + 1
		if len(utils.UrlArrays) == counter {
			fmt.Println("nytNewsletterCronJob: close chan", counter)
			close(c)
		}
	}
	fmt.Println("cronjobs/SetNytNewsLetterToRedis: channelStatus: close")

}

func getNewsAndSetData(key string, c chan NYTNewsCachingStruct) {
	result, err := network.PostCrawlGetNYTimeArrayEveningBriefing(key)
	if len(result) == 0 {
		fmt.Println("cronjobs/getNewsAndSetData", err)
	}
	c <- NYTNewsCachingStruct{
		eventType:        key,
		newsLetterStruct: result,
	}
}
