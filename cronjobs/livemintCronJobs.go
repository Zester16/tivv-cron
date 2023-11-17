package cronjobs

import (
	"encoding/json"
	"fmt"
	"stockpull/datasource"
	"stockpull/network"
	"stockpull/utils"

	"github.com/redis/go-redis/v9"
)

var RedisKeyMintNewLetter = "livemint-totm-nm"

func SetMintTopOfMorningNewsletter() {

	rdb := datasource.RedisConnect()

	newsletterString, redisErr := rdb.RedisDBConnector.Get(ctx, RedisKeyMintNewLetter).Result()

	newsBody := network.GetLiveMintNewsletter()

	newNewsArray := []NewsLetterStruct{{Date: utils.GetTodaysDateToString(), NewsBody: newsBody}}

	if redisErr != redis.Nil {
		oldNewsArray := []NewsLetterStruct{}
		err := json.Unmarshal([]byte(newsletterString), &oldNewsArray)
		if err != nil {
			fmt.Println(err)
		}

		if oldNewsArray[0].Date != newNewsArray[0].Date {
			newNewsArray = append(newNewsArray, oldNewsArray...)
		} else {
			newNewsArray = oldNewsArray
		}
	}

	j, _ := json.Marshal(newNewsArray)
	rdb.RedisDBConnector.Set(ctx, RedisKeyMintNewLetter, j, 0).Err()
}
