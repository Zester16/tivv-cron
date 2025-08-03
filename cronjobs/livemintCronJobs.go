package cronjobs

import (
	"encoding/json"
	"fmt"
	"stockpull/datasource"
	"stockpull/model"
	"stockpull/services"
	"stockpull/utils"

	"github.com/redis/go-redis/v9"
)

func SetMintTopOfMorningNewsletter() {
	fmt.Println("MinttopOfTheMorningCronJob")

	var RedisKeyMintNewLetter = model.MINT_TOP_OF_MORNING
	rdb := datasource.RedisConnect()

	newsletterString, redisErr := rdb.RedisDBConnector.Get(ctx, RedisKeyMintNewLetter).Result()

	url, urlError := services.GetLivemintTopOfTheDayUrl()

	if urlError != nil {
		fmt.Println("cron:SetMintTopOfMorningNewsletter err", urlError.Error())
	}

	//fmt.Println(newsBody)
	var newNewsArray = []NewsObject{{Date: utils.GetTodaysDateToString(), NewsUrl: url}}

	if redisErr != redis.Nil && urlError == nil {
		oldNewsArray := []NewsObject{}
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
	} else {
		fmt.Println("SetMintTopOfMorningNewsletter Error:", urlError)
	}

	j, err := json.Marshal(newNewsArray)
	if err != nil {
		fmt.Println("LivemintNewsMarshal Error: ", err)
	}
	rdb.RedisDBConnector.Set(ctx, RedisKeyMintNewLetter, j, 0).Err()
}
