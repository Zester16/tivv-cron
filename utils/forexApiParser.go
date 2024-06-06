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
	ZAR ForexData `json:"ZAR"`
	SGD ForexData `json:"SGD"`
	MXN ForexData `json:"MXN"`
	SEK ForexData `json:"SEK"`
	RUB ForexData `json:"RUB"`
	BRL ForexData `json:"BRL"`
	TRY ForexData `json:"TRY"`
	PLN ForexData `json:"PLN"`
	EGP ForexData `json:"EGP"`
	IDR ForexData `json:"IDR"`
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
		{Name: "SA Riyal", Code: forexObject.Data.SAR.Code, Value: forexObject.Data.SAR.Value},
		{Name: "ZA Rand", Code: forexObject.Data.ZAR.Code, Value: forexObject.Data.ZAR.Value},
		{Name: "Singapore Dollar", Code: forexObject.Data.SGD.Code, Value: forexObject.Data.SGD.Value},
		{Name: "Mexican Peso", Code: forexObject.Data.MXN.Code, Value: forexObject.Data.MXN.Value},
		{Name: "Swedish Krona", Code: forexObject.Data.SEK.Code, Value: forexObject.Data.SEK.Value},
		{Name: "Russian Ruble", Code: forexObject.Data.RUB.Code, Value: forexObject.Data.RUB.Value},
		{Name: "Brazilian Lira", Code: forexObject.Data.BRL.Code, Value: forexObject.Data.BRL.Value},
		{Name: "Turkish Liara", Code: forexObject.Data.TRY.Code, Value: forexObject.Data.TRY.Value},
		{Name: "Polish Zloty", Code: forexObject.Data.PLN.Code, Value: forexObject.Data.PLN.Value},
		{Name: "Eyptian Pound", Code: forexObject.Data.EGP.Code, Value: forexObject.Data.EGP.Value},
		{Name: "Indonesian Rupiya", Code: forexObject.Data.IDR.Code, Value: forexObject.Data.IDR.Value},
	}

	return currencyData
}
