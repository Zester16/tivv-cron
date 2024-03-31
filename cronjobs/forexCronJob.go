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

func SetForexCronJob() {
	const LOGGER = "cronjobs/SetForexCronJob/"
	newForexString := network.GetForex()

	rds := datasource.RedisConnect()

	resp, err := rds.RedisDBConnector.Get(ctx, model.FOREX_KEY_NAME).Result()

	var oldForexData = model.ForexRedisModel{}

	json.Unmarshal([]byte(resp), &oldForexData)

	newForexArray := []model.ForexIndividualModel{}
	if err == redis.Nil {
		for _, ele := range newForexString {
			newForexArray = append(newForexArray, model.ForexIndividualModel{Name: ele.Name, Code: ele.Code, Value: ele.Value, Change: 0, CP: 0})
		}
		fmt.Println(LOGGER, "S: ", "Updated data on", utils.GetTodaysDateToString())
	} else {
		for i, ele := range newForexString {
			oldVal := oldForexData.Data[i].Value
			change := ele.Value - oldVal
			cp := (change / oldVal) * 100
			newForexArray = append(newForexArray, model.ForexIndividualModel{Name: ele.Name, Code: ele.Code, Value: ele.Value, Change: change, CP: cp})
		}
		fmt.Println(LOGGER, "S: ", "Updated data by comparision with", oldForexData.Date)
	}

	newForexRedisObject := model.ForexRedisModel{
		Data: newForexArray,
		Date: utils.GetTodaysDateToString(),
	}

	nfro, _ := json.Marshal(newForexRedisObject)

	rds.RedisDBConnector.Set(ctx, model.FOREX_KEY_NAME, nfro, 0).Result()
}
