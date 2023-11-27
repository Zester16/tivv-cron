package controller

import (
	"net/http"
	"stockpull/cronjobs"
	"stockpull/datasource"
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

// func AddLiveMintNewsletterArray(w http.ResponseWriter, r *http.Request) {
// 	cronjobs.SetMintTopOfMorningNewsletter()
// 	w.Write([]byte("its all ok"))
// }
