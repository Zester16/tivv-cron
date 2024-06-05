package utils

import (
	"encoding/json"
)

func ParseLiveMintAllIndexJson(input string) []StockIndex {

	var dataArray []LiveMintStockStruct
	json.Unmarshal([]byte(input), &dataArray)

	var stockIndex []StockIndex
	for _, ind := range dataArray {
		stockIndex = append(stockIndex, StockIndex{StockIndexName: ind.Name, Points: ChangeFloattoString(ind.LivePrice), ChangePercent: ChangeFloattoString(ind.PercentChange)})
	}

	return stockIndex
}
