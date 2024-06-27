package cronjobs

import (
	"encoding/json"
	"fmt"
	"stockpull/datasource"
	"stockpull/model"
	"stockpull/repository"
	"stockpull/utils"
)

// gets stocks from respective network and then adds it in array
func SetAllStockCronJob() {
	const LOG_STRUCTURE = "cronjobs/SetAllStocksCronJob"
	rdb := datasource.RedisConnect()
	resp, _ := rdb.RedisDBConnector.Get(ctx, model.ALL_INDEX_KEY_NAME).Result()

	fmt.Println(LOG_STRUCTURE + resp)

	resStocks, err := repository.GetAllStockNews()

	if err != nil {
		fmt.Println(LOG_STRUCTURE, "/ERROR:", err.Error())
		return
	}

	stockIndexArray := model.StockIndexArray{
		Data: resStocks,
		Date: utils.GetTodaysDateToString(),
	}

	resultMarshalled, err := json.Marshal(stockIndexArray)

	if err != nil {
		fmt.Println(LOG_STRUCTURE+"/ERROR:", err.Error())
		return
	}

	rdb.RedisDBConnector.Set(ctx, model.ALL_INDEX_KEY_NAME, resultMarshalled, 0).Err()

}
