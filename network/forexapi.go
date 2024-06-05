package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"stockpull/utils"
)

func GetForex() []utils.CurrencyData {
	apiKey := os.Getenv("forex_api_key")
	url := "https://api.currencyapi.com/v3/latest?apikey=" + apiKey
	body, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	request, err := ioutil.ReadAll(body.Body)

	return utils.ForexApiJSONParser(string(request))

}
