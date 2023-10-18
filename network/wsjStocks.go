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

	var stockTickerArray = utils.ReadXML(sb)

	return stockTickerArray
}

func GetEuropeIndex() []utils.StockIndex {
	resp, err := http.Get("https://api.wsj.net/api/dylan/quotes/v2/comp/quoteByDialect?ckey=57494d5ed7&dialect=charting&id=INDEX%2FFR%2F%2FPX1%2CSX5P%2CINDEX%2FUK%2F%2FMCX&maxinstrumentmatches=1&needed=Meta%7CCompositeTrading%7CBlueGrassChannels&EntitlementToken=57494d5ed7ad44af85bc59a51dd87c90")
	if err != nil {
		fmt.Println(err)
	}
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
	usaIndexArray := GetWSJUSAIndex()

	var allIdex = append(europeIndexArray, asiaIndexAray...)
	allIdex = append(allIdex, usaIndexArray...)
	return allIdex
}

func GetWSJBonds() []utils.StockIndex {
	url := "https://api.wsj.net/api/dylan/quotes/v2/comp/quoteByDialect?ckey=57494d5ed7&dialect=charting&id=Bond%2FBX%2FXTUP%2FTMUBMUSD10Y%2CINDEX%2FUS%2FCBSX%2FVIX%2CGC00%2CFUTURE%2FUS%2F%2FCrude+Oil+-+Electronic%2CBUXX%2CBKX%2CSPGSCI&maxinstrumentmatches=1&needed=Meta%7CCompositeTrading%7CBlueGrassChannels&EntitlementToken=57494d5ed7ad44af85bc59a51dd87c90"

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	sb := string(body)
	fmt.Println(sb)
	list := utils.ReadXML(sb)

	return list
}
func GetWSJUSAIndex() []utils.StockIndex {
	url := "https://api.wsj.net/api/dylan/quotes/v2/comp/quoteByDialect?ckey=57494d5ed7&dialect=djid&id=343338%2C497001%2C343345%2C343303%2C433-25014677&maxinstrumentmatches=1&needed=Meta%7CCompositeTrading%7CBlueGrassChannels&EntitlementToken=57494d5ed7ad44af85bc59a51dd87c90"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	sb := string(body)
	list := utils.ReadXML(sb)

	return list
}
