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
	oldNewsArray := []NewsLetterStruct{}

	if redisErr != redis.Nil {
		if len(newsletterString) > 0 {
			err := json.Unmarshal([]byte(newsletterString), &oldNewsArray)

			if err != nil {
				fmt.Println(err)
			}
		}
		newNewsArray = append(newNewsArray, oldNewsArray...)
	}

	j, _ := json.Marshal(newNewsArray)
	rdb.RedisDBConnector.Set(ctx, RedisKeyMintNewLetter, j, 0).Err()
}
