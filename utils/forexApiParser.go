package utils

import (
	"encoding/json"
)

type CurrencyData struct {
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

type ForexHead struct {
	Data ForexObject `json:"data"`
}

type ForexObject struct {
	JPY ForexData `json:"JPY"`
	INR ForexData `json:"INR"`
	GBP ForexData `json:"GBP"`
	EUR ForexData `json:"EUR"`
	CNY ForexData `json:"CNY"`
	CHF ForexData `json:"CHF"`
	SAR ForexData `json:"SAR"`
	SGD ForexData `json:"SGD"`
}

type ForexData struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

func ForexApiJSONParser(input string) []CurrencyData {

	var forexObject ForexHead
	//fmt.Println(input)
	json.Unmarshal([]byte(input), &forexObject)

	//fmt.Println(forexObject.Data.INR)

	currencyData := []CurrencyData{{Name: "Indian Rupee", Code: forexObject.Data.INR.Code, Value: forexObject.Data.INR.Value},
		{Name: "Japanese Yen", Code: forexObject.Data.JPY.Code, Value: forexObject.Data.JPY.Value},
		{Name: "Pound", Code: forexObject.Data.GBP.Code, Value: forexObject.Data.GBP.Value},
		{Name: "Chinese Yuan", Code: forexObject.Data.CNY.Code, Value: forexObject.Data.CNY.Value},
		{Name: "Euro", Code: forexObject.Data.EUR.Code, Value: forexObject.Data.EUR.Value},
		{Name: "Swiss Franc", Code: forexObject.Data.CHF.Code, Value: forexObject.Data.CHF.Value},
		{Name: "SA Rand", Code: forexObject.Data.SAR.Code, Value: forexObject.Data.SAR.Value},
		{Name: "Singapore Dollar", Code: forexObject.Data.SGD.Code, Value: forexObject.Data.SGD.Value}}

	return currencyData
}
