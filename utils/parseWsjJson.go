package utils

import (
	"encoding/json"
	"fmt"
)

type WSJData struct {
	Datas []WSJJson `json:"data"`
}

type WSJJson struct {
	CommonName string `json:"commonName"`
	Value      string `json:"value"`
	PerChange  string `json:"perChange"`
}

func ParseWsjJson(input string) []StockIndex {

	var dataArray WSJData

	err := json.Unmarshal([]byte(input), &dataArray)

	if err != nil {
		fmt.Println("WSJ-JSON", err)
	}
	var stockIndexArray []StockIndex

	for i := 0; i < len(dataArray.Datas); i++ {
		stockIndex := StockIndex{StockIndexName: dataArray.Datas[i].CommonName, ChangePercent: dataArray.Datas[i].PerChange, Points: dataArray.Datas[i].Value}
		stockIndexArray = append(stockIndexArray, stockIndex)
	}

	return stockIndexArray

}
