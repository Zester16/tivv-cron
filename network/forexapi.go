package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"stockpull/utils"
)

func GetForex() []utils.CurrencyData {
	url := "https://api.currencyapi.com/v3/latest?apikey=cur_live_M1vQQ90ktOOaoCpb290qIYSxOTFsEKwkk0jtXq6J"
	body, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	request, err := ioutil.ReadAll(body.Body)

	return utils.ForexApiJSONParser(string(request))

}
