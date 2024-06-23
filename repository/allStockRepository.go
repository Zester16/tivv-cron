package repository

import (
	"stockpull/network"
	"stockpull/utils"
)

//This repository deals with pulling stock market news from all sources from network and pass to respective requestor

func GetAllStockNews() ([]utils.StockIndex, error) {

	wsjStockIndex := network.GetAllIndex()
	indianStockIndex, err := network.GetMintLiveAllIndex()

	if err != nil {
		return []utils.StockIndex{}, err
	}

	newIndex := []utils.StockIndex{}
	newIndex = append(newIndex, wsjStockIndex...)
	newIndex = append(newIndex, indianStockIndex...)

	return newIndex, nil
}
