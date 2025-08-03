package controller

import (
	"encoding/json"
	"net/http"
	"stockpull/datasource"
	"stockpull/model"
	"stockpull/network"
	"stockpull/repository"
	"stockpull/services"
	"stockpull/utils"
)

func GetLiveMintNewsletterArray(w http.ResponseWriter, r *http.Request) {
	// rdb := datasource.RedisConnect()

	// lmNewsArray, rdbError := rdb.RedisDBConnector.Get(ctx, cronjobs.RedisKeyMintNewLetter).Result()

	// if rdbError != nil {

	// 	w.Write([]byte(rdbError.Error()))
	// }
	//cronjobs.SetMintTopOfMorningNewsletter()

	//w.Header().Add("Content-Type", "application/json")
	url, err := services.GetLivemintTopOfTheDayUrl()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(url))
}

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
