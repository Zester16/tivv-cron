package controller

import (
	"encoding/json"
	"net/http"
	"stockpull/network"
)

// controller handles REST endpoint relating to WSJ
func Test(w http.ResponseWriter, r *http.Request) {
	sb := network.GetAsiaIndex()
	j, _ := json.Marshal(sb)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func WsjAsia(w http.ResponseWriter, r *http.Request) {
	sb := network.GetAsiaIndex()
	j, _ := json.Marshal(sb)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
