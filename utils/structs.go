package utils

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
	Change         string `json:"change"`
}

// struct specific to json
type LiveMintStockStruct struct {
	ExchangeType  string  `json:"exchangeType"`
	Name          string  `json:"name"`
	LivePrice     float64 `json:"livePrice"`
	PercentChange float64 `json:"percentChange"`
}
