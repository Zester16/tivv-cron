package controller

import (
	"net/http"
	"stockpull/datasource"
	"stockpull/model"
	"stockpull/utils"
)

// function gets an array of globe and mail newsletters
func GetGlobeAndMailNewsletters(w http.ResponseWriter, r *http.Request) {

	rdb := datasource.RedisConnect()

	globeMailNewsArray, rdbError := rdb.RedisDBConnector.Get(ctx, model.GLOBE_MAIL).Result()

	if rdbError != nil {
		utils.SetError400(w, rdbError)
	}

	w.Header().Add("Content-Type", "Application/json")
	w.Write([]byte(globeMailNewsArray))

}
