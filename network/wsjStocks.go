package network

import (
	"fmt"
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

func GetEuropeIndex() []utils.StockIndex {
	resp, err := http.Get("https://api.wsj.net/api/dylan/quotes/v2/comp/quoteByDialect?ckey=57494d5ed7&dialect=charting&id=INDEX%2FFR%2F%2FPX1%2CSX5P%2CINDEX%2FUK%2F%2FMCX&maxinstrumentmatches=1&needed=Meta%7CCompositeTrading%7CBlueGrassChannels&EntitlementToken=57494d5ed7ad44af85bc59a51dd87c90")
	respNew, err := http.Get("https://api.wsj.net/api/dylan/quotes/v2/comp/quoteByDialect?ckey=57494d5ed7&dialect=djid&id=343341%2C440104%2C463275%2C353778&maxinstrumentmatches=1&needed=Meta%7CCompositeTrading%7CBlueGrassChannels&EntitlementToken=57494d5ed7ad44af85bc59a51dd87c90")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	newBody, err := ioutil.ReadAll(respNew.Body)

	sb := string(body)
	sbTwo := string(newBody)
	var stockTickerArray = utils.ReadXML(sb)
	var stockTickerArrayTwo = utils.ReadXML(sbTwo)
	stockTickerArray = append(stockTickerArray, stockTickerArrayTwo...)
	return stockTickerArray
}

func GetAllIndex() []utils.StockIndex {
	europeIndexArray := GetEuropeIndex()
	asiaIndexAray := GetAsiaIndex()

	var allIdex = append(europeIndexArray, asiaIndexAray...)
	return allIdex
}
