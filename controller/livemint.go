package controller

import (
	"encoding/json"
	"net/http"
	"stockpull/datasource"
	"stockpull/model"
	"stockpull/network"
	"stockpull/repository"
	"stockpull/utils"
)

/*
* To get an array of livemint top of the morning news
 */
func GetLiveMintNewsletterArray(w http.ResponseWriter, r *http.Request) {
	rdb := datasource.RedisConnect()

	lmNewsArray, rdbError := rdb.RedisDBConnector.Get(ctx, model.MINT_TOP_OF_MORNING).Result()

	if rdbError != nil {

		w.Write([]byte(rdbError.Error()))
	}
	//for live newsletter
	//cronjobs.SetMintTopOfMorningNewsletter()
	//url, err := services.GetLivemintTopOfTheDayUrl()
	w.Header().Add("Content-Type", "application/json")

	if rdbError != nil {
		w.Write([]byte(rdbError.Error()))
		return
	}

	w.Write([]byte(lmNewsArray))
}

// to get all 3 stock market indexes live than morning cached data
func GetMintLiveNewsArray(w http.ResponseWriter, r *http.Request) {
	resp, err := network.GetMintLiveAllIndex()

	if err != nil {
		utils.SetError400(w, err)
	}

	w.Header().Add("Content-Type", "application/json")
	strResp, err := json.Marshal(resp)

	if err != nil {
		utils.SetError400(w, err)
	}
	w.Write(strResp)
}

// gets stock all live stock index from WSJ and Livemint and sends as one
func GetAllStockMarkets(w http.ResponseWriter, r *http.Request) {

	resp, err := repository.GetAllStockNews()

	if err != nil {
		utils.SetError400(w, err)
	}

	respStr, err := json.Marshal(resp)

	if err != nil {
		utils.SetError400(w, err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(respStr)
}

// gets all cached stock market indexes which runs in morning
func GetAllStocksCached(w http.ResponseWriter, r *http.Request) {
	rdx := datasource.RedisConnect()

	resp, _ := rdx.RedisDBConnector.Get(ctx, model.ALL_INDEX_KEY_NAME).Result()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp))
}

// func AddLiveMintNewsletterArray(w http.ResponseWriter, r *http.Request) {
// 	cronjobs.SetMintTopOfMorningNewsletter()
// 	w.Write([]byte("its all ok"))
// }
