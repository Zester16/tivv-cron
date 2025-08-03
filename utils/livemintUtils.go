package utils

import (
	"encoding/json"
	"stockpull/model"
)

// converter function
// gets stock index from livemint and then does data transformation on index to suite all indexes
func ParseLiveMintAllIndexJson(input string) []model.StockIndex {

	var dataArray []model.LiveMintStockStruct
	json.Unmarshal([]byte(input), &dataArray)

	var stockIndex []model.StockIndex
	for _, ind := range dataArray {
		stockIndex = append(stockIndex, model.StockIndex{StockIndexName: ind.Name, Points: ChangeFloattoString(ind.LivePrice), ChangePercent: ChangeFloattoString(ind.PercentChange), ChangePoint: ChangeFloattoString(ind.NetChange)})
	}

	return stockIndex
}
