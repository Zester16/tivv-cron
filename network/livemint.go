package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"stockpull/utils"
)

func GetLiveMintNewsletter() string {

	url := "https://www.livemint.com/mint-top-newsletter/minttopofthemorning.html"
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	reqBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	db := string(reqBody)
	return db

}

func GetMintLiveAllIndex() ([]utils.StockIndex, error) {
	URL := "https://api-mintgenie.livemint.com/api-gateway/fundamental/api/v2/indices/home/getHomeIndices?forMarkets=false"

	resp, err := http.Get(URL)

	if err != nil {
		return []utils.StockIndex{}, err
	}

	parsedBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []utils.StockIndex{}, err
	}

	parsedArray := utils.ParseLiveMintAllIndexJson(string(parsedBody))

	return parsedArray, nil
}
