package cronjobs

import (
	"encoding/json"
	"fmt"
	"stockpull/datasource"
	"stockpull/model"
	"stockpull/network"
	"stockpull/utils"

	"github.com/redis/go-redis/v9"
)

func BmlCronJob() {
	const LOG_STRUCTURE = "cronjobs/BlmCronJobs"
	rdb := datasource.RedisConnect()
	blmUrls := model.BlmTest.GetBLMUrls()

	for key, value := range blmUrls {

		response, err := network.PostCrawlGetBloombergNewsLetter(value)

		if err != nil {
			fmt.Println(LOG_STRUCTURE, "E: ", err)
			fmt.Println(LOG_STRUCTURE, err.Error())
			return
		}

		blmData, err := rdb.RedisDBConnector.Get(ctx, key).Result()
		newsArray := []NewsObject{{Date: utils.GetTodaysDateToString(),
			NewsUrl: response}}

		if err != redis.Nil {
			oldNewsArray := []NewsObject{}
			json.Unmarshal([]byte(blmData), &oldNewsArray)

			if oldNewsArray[0].NewsUrl == response {
				newsArray = oldNewsArray
			} else {
				newsArray = append(newsArray, oldNewsArray...)
			}

		}

		j, _ := json.Marshal(newsArray)
		rdb.RedisDBConnector.Set(ctx, key, j, 0).Err()

	}

}
