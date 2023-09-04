package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetBloombergTechNews(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	bloomberg := "https://www.bloomberg.com/newsletters/five-things-europe/latest"

	req, _ := http.NewRequest("GET", bloomberg, nil)
	req.Header = http.Header{"Cookie": {"table: 0x7f0fe1d04bf8; exp_pref=AMER; seen_uk=1; country_code=US"}, "Host": {"www.bloomberg.com"}, "UserAgent": {"PostmanRuntime/7.32.3"}}
	//req.Header.Set()

	rep, _ := client.Do(req)

	//body, err := http.Get(bloomberg)

	fmt.Print(rep.StatusCode)
	response, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		//log.Fatal(err)
		//w.Write([]byte("Not our page"))
	}
	w.Write(response)

}
