package model

import "stockpull/utils"

type StockIndexArray struct {
	Data []utils.StockIndex `json:"data"`
	Date string        `json:"date"`
}

const ALL_INDEX_KEY_NAME = "all_country_stock_index_mh"
