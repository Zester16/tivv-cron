package repository

import (
	"stockpull/model"
	"stockpull/network"
)

//This repository deals with pulling stock market news from all sources from network and pass to respective requestor

func GetAllStockNews() ([]model.StockIndex, error) {

	wsjStockIndex := network.GetAllIndex()
	indianStockIndex, err := network.GetMintLiveAllIndex()

	if err != nil {
		return []model.StockIndex{}, err
	}

	newIndex := []model.StockIndex{}
	newIndex = append(newIndex, wsjStockIndex...)
	newIndex = append(newIndex, indianStockIndex...)

	return newIndex, nil
}
