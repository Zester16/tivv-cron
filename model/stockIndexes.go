package model

type StockIndexArray struct {
	Data []StockIndex `json:"data"`
	Date string       `json:"date"`
}

const ALL_INDEX_KEY_NAME = "all_country_stock_index_mh"
