package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetBloombergTechNews(w http.ResponseWriter, r *http.Request) {
	//bloomberg := "https://www.bloomberg.com/newsletters/five-things-europe/lates"
	body, err := http.Get("https://m.com")
	if err != nil {
		//log.Fatal(err)
		//w.Write([]byte("Not our page"))
	}
	fmt.Print(body.StatusCode)
	response, err := ioutil.ReadAll(body.Body)
	w.Write(response)

}
