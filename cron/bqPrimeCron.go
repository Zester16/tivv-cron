package cron

import (
	"context"
	"encoding/json"
	"stockpull/datasource"
	"stockpull/utils"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type NewsObject struct {
	Date    string `json:"date"`
	NewsUrl string `json:"newsUrl"`
}

func SetBqPrimeNEwsLetterArray() {
	bqPrimeName := "bqprimeArray"
	rdb := datasource.RedisConnect()

	bqArrayString, err := rdb.RedisDBConnector.Get(ctx, bqPrimeName).Result()

	url := utils.GetBQPrimeUrl()

	var bqArray = []NewsObject{{Date: "23-09-2023", NewsUrl: url}}

	if err != redis.Nil {
		newsObject := []NewsObject{}
		err := json.Unmarshal([]byte(bqArrayString), &newsObject)
		bqArray = append(bqArray, newsObject...)
		if err != nil {
			bqArray = append(bqArray, newsObject...)
		}

	}
	j, _ := json.Marshal(bqArray)
	rdb.RedisDBConnector.Set(ctx, bqPrimeName, j, 0).Err()

}
