package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
