package model

//GENERIC STRUCTS
type NewsLetterNytStruct struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Img   string `json:"img"`
	Date  string `json:"date"`
}
type StockIndex struct {
	StockIndexName string `json:"stockIndexName"`
	Points         string `json:"points"`
	ChangePercent  string `json:"changePercent"`
	ChangePoint    string `json:"changePoint"`
}

// struct specific to json
type LiveMintStockStruct struct {
	ExchangeType  string  `json:"exchangeType"`
	Name          string  `json:"name"`
	LivePrice     float64 `json:"livePrice"`
	PercentChange float64 `json:"percentChange"`
	NetChange     float64 `json:"netChange"`
}
