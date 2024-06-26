package cronjobs

import (
	"context"
	"encoding/json"
	"fmt"
	"stockpull/datasource"
	"stockpull/utils"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var BqPrimeName = "bqprimeArray"

func SetBqPrimeNEwsLetterArray() {

	rdb := datasource.RedisConnect()

	bqArrayString, err := rdb.RedisDBConnector.Get(ctx, BqPrimeName).Result()

	url := utils.GetBQPrimeUrl()

	fmt.Println("bqPrimeCron: ", url)

	tm := time.Now()
	month := tm.Month().String()
	day := strconv.Itoa(tm.Day())
	yr := strconv.Itoa(tm.Year())
	todaysDate := day + "-" + month + "-" + yr

	var bqArray = []NewsObject{{Date: todaysDate, NewsUrl: url}}

	if err != redis.Nil {
		newsObject := []NewsObject{}
		errJsn := json.Unmarshal([]byte(bqArrayString), &newsObject)

		isWeekend := utils.CheckTodayIsWeekend()

		if bqArray[0].Date != newsObject[0].Date && !isWeekend {
			bqArray = append(bqArray, newsObject...)

		} else {
			bqArray = newsObject
		}

		fmt.Println(newsObject[0].Date)
		if errJsn != nil {
			fmt.Printf(errJsn.Error())
		}

	}
	j, _ := json.Marshal(bqArray)
	rdb.RedisDBConnector.Set(ctx, BqPrimeName, j, 0).Err()

}
