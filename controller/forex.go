package controller

import (
	"encoding/json"
	"net/http"
	"stockpull/network"
)

func GetCurrencyValue(w http.ResponseWriter, r *http.Request) {

	resp := network.GetForex()

	j, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(j))

}
