package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"stockpull/datasource"
	"stockpull/model"
	"stockpull/utils"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

/*TODO: MOVE ALL REDIS GET AND POST LOGIC TO REPOSITORY*/
/*
* wil hit total part of database. the set present in model will loop through and pull data from each of list and display the news
 */
func GetTotalNewsletterCount() (model.TotalNewslettersArray, error) {
	funcName := "repository.GetTotalNewsletterCount :"
	rdb := datasource.RedisConnect()
	listStr, err := rdb.RedisDBConnector.Get(ctx, model.TOTAL).Result()
	var list model.TotalNewslettersArray
	if err == redis.Nil {
		obj, err := createTotalNewsLetterListAndSetinDB()

		if err != nil {
			fmt.Println(funcName, "ERROR :", err.Error())
		}
		return obj, err
	} else if err != nil {
		if err != nil {
			fmt.Println(funcName, "DB ERROR :", err.Error())
		}
		return list, err
	} else {
		err := json.Unmarshal([]byte(listStr), &list)
		if err != nil {
			fmt.Println(funcName, "Json ERROR :", err.Error())
		}
	}
	return list, nil

}

func createTotalNewsLetterListAndSetinDB() (model.TotalNewslettersArray, error) {
	funcName := "repository.createTotalNewsLetterListAndSetinDB"
	tnla := model.TotalNewslettersArray{Date: utils.GetTodaysDateToString(), Data: []model.PerNewsLetterCountStruct{}}
	rdb := datasource.RedisConnect()
	for key, va := range model.RedisKey {
		fmt.Println(key, va)
		listStr, err := rdb.RedisDBConnector.Get(ctx, va).Result()

		if err == nil {
			var list []model.NewsLetterDate

			json.Unmarshal([]byte(listStr), &list)
			//fmt.Println(funcName, list)
			elements := len(list)

			tnla.Data = append(tnla.Data, model.PerNewsLetterCountStruct{Name: key, Total: elements})

		} else {
			fmt.Println(funcName, "Data Pull ERROR: ", err.Error())
			//return tnla, err
		}

	}
	tnlaStr, err := json.Marshal(tnla)

	if err != nil {
		fmt.Println(funcName, " JSON ERROR: ", err.Error())
		return tnla, err
	}

	now := time.Now()
	tomorrowsMorning := time.Date(now.Year(), now.Month(), now.Day()+1, 2, 45, 0, 0, time.UTC)

	rdb.RedisDBConnector.Set(ctx, model.TOTAL, tnlaStr, 0)
	rdb.RedisDBConnector.ExpireAt(ctx, model.TOTAL, tomorrowsMorning)

	return tnla, nil
}
