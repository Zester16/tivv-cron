package controller

import (
	"encoding/json"
	"net/http"
	"stockpull/cronjobs"
	"stockpull/datasource"
	"stockpull/network"
	"stockpull/utils"
)

func GetLiveMintNewsletterArray(w http.ResponseWriter, r *http.Request) {
	rdb := datasource.RedisConnect()

	lmNewsArray, rdbError := rdb.RedisDBConnector.Get(ctx, cronjobs.RedisKeyMintNewLetter).Result()

	if rdbError != nil {

		w.Write([]byte(rdbError.Error()))
	}

	//cronjobs.SetMintTopOfMorningNewsletter()

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(lmNewsArray))
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

// func AddLiveMintNewsletterArray(w http.ResponseWriter, r *http.Request) {
// 	cronjobs.SetMintTopOfMorningNewsletter()
// 	w.Write([]byte("its all ok"))
// }
