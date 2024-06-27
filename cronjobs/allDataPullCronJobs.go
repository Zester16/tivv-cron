package cronjobs

import (
	"encoding/json"
	"fmt"
	"stockpull/datasource"
	"stockpull/model"
	"stockpull/repository"
	"stockpull/utils"

	"github.com/redis/go-redis/v9"
)

// gets stocks from respective network and then adds it in array
func SetAllStockCronJob() {
	const LOG_STRUCTURE = "cronjobs/SetAllStocksCronJob"
	rdb := datasource.RedisConnect()
	_, err := rdb.RedisDBConnector.Get(ctx, model.ALL_INDEX_KEY_NAME).Result()

	if err != redis.Nil {
		fmt.Println(LOG_STRUCTURE, "/ERROR:", err.Error())
		return
	}

	resStocks, err := repository.GetAllStockNews()

	stockIndexArray := model.StockIndexArray{
		Data: resStocks,
		Date: utils.GetTodaysDateToString(),
	}

	resultMarshalled, err := json.Marshal(stockIndexArray)

	if err != nil {
		fmt.Println(LOG_STRUCTURE, "/ERROR:", err.Error())
		return
	}

	rdb.RedisDBConnector.Set(ctx, model.ALL_INDEX_KEY_NAME, resultMarshalled, 0).Err()

}
