package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"stockpull/model"
	"stockpull/utils"
)

// passes livmint url to check if url works or not
func CheckLiveMintNewsletterUrl(url string) bool {

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error: CheckLiveMintNewsletterUrl:", err)
		return false
	}

	if response.StatusCode != 200 {
		fmt.Println("Error: CheckLiveMintNewsletterUrl:", response.StatusCode)
		return false
	}

	return true

}

func GetMintLiveAllIndex() ([]model.StockIndex, error) {
	URL := "https://api-mintgenie.livemint.com/api-gateway/fundamental/api/v2/indices/home/getHomeIndices?forMarkets=false"

	resp, err := http.Get(URL)

	if err != nil {
		return []model.StockIndex{}, err
	}

	parsedBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []model.StockIndex{}, err
	}

	parsedArray := utils.ParseLiveMintAllIndexJson(string(parsedBody))

	return parsedArray, nil
}
