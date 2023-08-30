package network

import (
	"io/ioutil"
	"log"
	"net/http"
	"stockpull/utils"
)

func GetAsiaIndex() []utils.StockIndex {

	resp, err := http.Get("https://api.wsj.net/api/dylan/quotes/v2/comp/quoteByDialect?ckey=57494d5ed7&dialect=charting&id=INDEX%2FXX%2F%2FADOW%2CINDEX%2FHK%2FXHKG%2FHSI%2CINDEX%2FSG%2FXSES%2FSTI%2C%2FKR%2F%2FSEU&maxinstrumentmatches=1&needed=Meta%7CCompositeTrading%7CBlueGrassChannels&EntitlementToken=57494d5ed7ad44af85bc59a51dd87c90")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	sb := string(body)

	log.Printf(sb)
	var stockTickerArray = utils.ReadXML(sb)

	return stockTickerArray
}
