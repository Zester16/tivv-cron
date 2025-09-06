package controller

import (
	"encoding/json"
	"net/http"
	"stockpull/repository"
	"stockpull/utils"
)

// returns back an object with data having array of news names, each having count of how many newsletters are present
func GetNewsletterCount(w http.ResponseWriter, r *http.Request) {
	resp, err := repository.GetTotalNewsletterCount()

	if err != nil {
		utils.SetError400(w, err)
		return
	}

	respStr, err := json.Marshal(resp)
	if err != nil {
		utils.SetError400(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(respStr)
}
