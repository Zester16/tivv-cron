package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"stockpull/cronjobs"
	"stockpull/datasource"
	"stockpull/model"
	"stockpull/utils"
)

var ctx = context.Background()

func GetBloombergTechNews(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	//bloomberg := "https://www.bloomberg.com/newsletters/five-things-europe/latest"
	bq := "https://www.bqprime.com/markets/stock-market-today-all-you-need-to-know-going-into-trade-on-aug-07-02"
	req, _ := http.NewRequest("GET", bq, nil)
	//req.Header = http.Header{"Cookie": {"table: 0x7f0fe1d04bf8; exp_pref=AMER; seen_uk=1; country_code=US"}, "Host": {"www.bloomberg.com"}, "User-Agent": {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36"}}
	//req.Header.Set("User-Access", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36")
	//req.Header.Set("Cookie", "table: 0x7f0fe1d04bf8; exp_pref=AMER; seen_uk=1; country_code=US")
	rep, _ := client.Do(req)

	//body, err := http.Get(bloomberg)

	fmt.Print(rep.StatusCode)
	response, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		//log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Not our page"))
	}
	rep.Body.Close()
	w.Write(response)

}

// gets bloomberg news
// sets header and performs operation
type ResponseUrl struct {
	Url string `json:"url"`
}

// endpoint to get single day news letter directly
func GetBQPrimeTodaysAllYouNeedToKnowNews(w http.ResponseWriter, r *http.Request) {
	respObject := ResponseUrl{Url: utils.GetBQPrimeUrl()}

	//fmt.Print(fullString)
	//w.Write([]byte(string({"url":GetBQPrimeUrl()))})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respObject)
}

// endpoint to get array of newsletters of bqprime morning newsletter
func GetBQPrimeAllYouNeedToKnowArray(w http.ResponseWriter, r *http.Request) {
	//cron.SetBqPrimeNEwsLetterArray()
	rdb := datasource.RedisConnect()
	list, err := rdb.RedisDBConnector.Get(ctx, cronjobs.BqPrimeName).Result()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(list))

}

func GetBloombergEconomicsNewsLetter(w http.ResponseWriter, r *http.Request) {
	rdb := datasource.RedisConnect()
	resArr, err := rdb.RedisDBConnector.Get(ctx, model.BLM_ECO).Result()

	if err != nil {
		fmt.Println("", err.Error())
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))

	}
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(resArr))

	//	hitUrl := os.Getenv("blm_eco")
	//	response, err := network.PostCrawlGetBloombergNewsLetter(hitUrl)

	// w.Header().Set("Content-Type", "application/json")
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	w.Write([]byte(err.Error()))
	// }

	// w.Write([]byte(response))

}

// func GetLiveMintNewsletterArray(w http.ResponseWriter, r *http.Request) {
// 	//	rdb := datasource.RedisConnect()
// 	cronjobs.SetMintTopOfMorningNewsletter()

// 	w.Write([]byte("Its done"))
// }
