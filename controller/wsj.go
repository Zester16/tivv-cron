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

func WsjEurope(w http.ResponseWriter, r *http.Request) {
	sb := network.GetEuropeIndex()
	j, _ := json.Marshal(sb)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func WSJAllIndex(w http.ResponseWriter, r *http.Request) {
	sb := network.GetAllIndex()
	j, _ := json.Marshal(sb)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(j)
}

func WSJBonds(w http.ResponseWriter, r *http.Request) {
	sb := network.GetWSJBonds()
	j, _ := json.Marshal(sb)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(j)
}

func WSJUsaIndex(w http.ResponseWriter, r *http.Request) {
	sb := network.GetWSJUSAIndex()
	j, _ := json.Marshal(sb)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(j)
}
