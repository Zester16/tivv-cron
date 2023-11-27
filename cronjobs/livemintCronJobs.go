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
	fmt.Println("MinttopOfTheMorningCronJob")
	rdb := datasource.RedisConnect()

	newsletterString, redisErr := rdb.RedisDBConnector.Get(ctx, RedisKeyMintNewLetter).Result()

	newsBody := network.GetLiveMintNewsletter()
	//fmt.Println(newsBody)

	newNewsArray := []NewsLetterStruct{{Date: utils.GetTodaysDateToString(), NewsBody: newsBody}}

	if redisErr != redis.Nil {
		oldNewsArray := []NewsLetterStruct{}
		err := json.Unmarshal([]byte(newsletterString), &oldNewsArray)
		if err != nil {
			fmt.Println(err)
		}
		isWeekend := utils.CheckTodayIsWeekend()
		fmt.Println(isWeekend)
		if (oldNewsArray[0].Date != newNewsArray[0].Date) && !isWeekend {
			newNewsArray = append(newNewsArray, oldNewsArray...)
			fmt.Println("LivemintNewsCronJob", "Its weekday")
		} else {
			newNewsArray = oldNewsArray
			fmt.Println("LivemintNewsCronJob", "Its weekend")
		}
	}

	j, err := json.Marshal(newNewsArray)
	if err != nil {
		fmt.Println("LivemintNewsMarshal Error: ", err)
	}
	rdb.RedisDBConnector.Set(ctx, RedisKeyMintNewLetter, j, 0).Err()
}
