package controller

import (
	"encoding/json"
	"net/http"
	"stockpull/datasource"
	"stockpull/model"
	"stockpull/network"
)

func GetCurrencyValue(w http.ResponseWriter, r *http.Request) {

	resp := network.GetForex()

	j, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(j))

}

func GetCachedCurrencyValue(w http.ResponseWriter, r *http.Request) {

	rdx := datasource.RedisConnect()

	resp, _ := rdx.RedisDBConnector.Get(ctx, model.FOREX_KEY_NAME).Result()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp))

}
